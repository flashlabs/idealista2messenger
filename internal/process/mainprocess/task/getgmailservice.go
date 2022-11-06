package task

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"net/http"
)

func GetGmailService(client *http.Client) (*gmail.Service, error) {
	srv, err := gmail.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to retrieve Gmail client: %v", err))
	}

	return srv, nil
}
