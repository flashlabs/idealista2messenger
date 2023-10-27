package task

import (
	"fmt"

	"google.golang.org/api/gmail/v1"
)

// FetchMessages fetches user messages based on query and returns them.
func FetchMessages(srv *gmail.Service, userID, query string) (*gmail.ListMessagesResponse, error) {
	r, err := srv.Users.Messages.List(userID).Q(query).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve messages: %w", err)
	}

	return r, nil
}
