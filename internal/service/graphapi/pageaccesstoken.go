package graphapi

import (
	"encoding/json"
	"fmt"
	"os"
)

type PageAccessToken struct {
	Token string `json:"page_access_token"`
}

func PageAccessTokenFromFile(file string) (*PageAccessToken, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error while opening page access token file: %w", err)
	}

	defer f.Close()

	tok := &PageAccessToken{}
	err = json.NewDecoder(f).Decode(tok)

	if err != nil {
		return nil, fmt.Errorf("error while decoding page access token file: %w", err)
	}

	return tok, nil
}
