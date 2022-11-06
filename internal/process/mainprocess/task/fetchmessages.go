package task

import (
	"errors"
	"fmt"
	"google.golang.org/api/gmail/v1"
)

func FetchMessages(srv *gmail.Service) (*gmail.ListMessagesResponse, error) {
	user := "me"
	query := "is:unread"
	r, err := srv.Users.Messages.List(user).Q(query).Do()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to retrieve messages: %v", err))
	}

	return r, err
}
