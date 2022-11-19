package task

import (
	"context"
	"errors"
	"fmt"
	"github.com/flashlabs/idealista2messenger/internal/token"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"net/http"
)

func GetClient(credentials []byte) (*http.Client, error) {
	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(credentials, gmail.GmailModifyScope)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to parse client secret file to config: %v", err))
	}
	return getClient(config), nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "config/token.json"
	tok, err := token.AccessTokenFromFile(tokFile)
	if err != nil {
		tok = token.GetTokenFromWeb(config)
		token.SaveAccessToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}
