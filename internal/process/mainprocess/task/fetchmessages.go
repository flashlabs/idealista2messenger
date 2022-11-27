package task

import (
	"errors"
	"fmt"
	"google.golang.org/api/gmail/v1"
)

func FetchMessages(srv *gmail.Service, userId, query string) (*gmail.ListMessagesResponse, error) {
	r, err := srv.Users.Messages.List(userId).Q(query).Do()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to retrieve messages: %v", err))
	}

	return r, err
}
