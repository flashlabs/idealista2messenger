package task

import (
	"google.golang.org/api/gmail/v1"
)

func MarkRead(srv *gmail.Service, r *gmail.ListMessagesResponse) error {
	user := "me"
	for _, m := range r.Messages {
		modifyMessageRequest := gmail.ModifyMessageRequest{
			RemoveLabelIds: []string{"UNREAD"},
		}
		_, err := srv.Users.Messages.Modify(user, m.Id, &modifyMessageRequest).Do()
		if err != nil {
			return err
		}
	}

	return nil
}
