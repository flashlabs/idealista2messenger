package task

import (
	"context"
	"fmt"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"net/http"
)

func GmailService(client *http.Client) (*gmail.Service, error) {
	srv, err := gmail.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Gmail client: %v", err)
	}

	return srv, nil
}
