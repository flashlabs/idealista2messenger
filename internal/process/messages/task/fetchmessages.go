package task

import (
	"fmt"

	"google.golang.org/api/gmail/v1"
)

// FetchMessages fetches user messages based on query and returns them
func FetchMessages(srv *gmail.Service, userId, query string) (*gmail.ListMessagesResponse, error) {
	r, err := srv.Users.Messages.List(userId).Q(query).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve messages: %v", err)
	}

	return r, err
}
