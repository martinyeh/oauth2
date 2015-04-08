package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
)

var (
	clientId     = flag.String("clientid", "742969766528-rcpf31ri5jm3i9gmbk7iqthtkid63itg.apps.googleusercontent.com", "app's client id")
	clientSecret = flag.String("secret", "xgXHnF5AGb6oQCPQHXcAWi7q", "app's client secret")
)

var conf = &oauth2.Config{
	ClientID:     "742969766528-rcpf31ri5jm3i9gmbk7iqthtkid63itg.apps.googleusercontent.com", // Set by --clientid or --clientid_file
	ClientSecret: "xgXHnF5AGb6oQCPQHXcAWi7q",                                                 // Set by --secret or --secret_file
	RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
	Scopes:       []string{youtube.YoutubeScope}, // filled in per-API
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://accounts.google.com/o/oauth2/auth",
		TokenURL: "https://accounts.google.com/o/oauth2/token",
	},
}

func TokenFromWeb() *http.Client {
	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v\n", url)

	// Use the authorization code that is pushed to the redirect URL.
	// NewTransportWithCode will do the handshake to retrieve
	// an access token and initiate a Transport that is
	// authorized and authenticated by the retrieved token.
	var code string
	fmt.Printf("Please enter auth code:")
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("token:", tok)

	client := conf.Client(oauth2.NoContext, tok)

	return client
}

func main() {
	flag.Parse()
	conf.ClientID = *clientId
	conf.ClientSecret = *clientSecret

	log.Println("oauth2 conf:", conf)

	TokenFromWeb()
}
