package graphapi

import (
	"encoding/json"
	"os"
)

type PageAccessToken struct {
	Token string `json:"page_access_token"`
}

func PageAccessTokenFromFile(file string) (*PageAccessToken, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &PageAccessToken{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}
