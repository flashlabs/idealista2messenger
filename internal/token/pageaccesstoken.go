package token

import (
	"encoding/json"
	"github.com/flashlabs/idealista2messenger/internal/structs"
	"os"
)

func PageAccessTokenFromFile(file string) (*structs.PageAccessToken, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &structs.PageAccessToken{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}
