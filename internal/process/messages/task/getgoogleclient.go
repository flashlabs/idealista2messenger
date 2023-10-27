package task

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"

	"github.com/flashlabs/idealista2messenger/internal/token"
)

// GoogleClient creates new Google Client and returns it.
func GoogleClient(credentials []byte, tokenFile string) (*http.Client, error) {
	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(credentials, gmail.GmailModifyScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %w", err)
	}

	return client(config, tokenFile), nil
}

// Retrieve a token, saves the token, then returns the generated client.
func client(config *oauth2.Config, tokenFile string) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tok, err := token.AccessTokenFromFile(tokenFile)
	if err != nil {
		tok = token.AccessTokenFromWeb(config)
		token.SaveAccessToken(tokenFile, tok)
	}

	return config.Client(context.Background(), tok)
}
