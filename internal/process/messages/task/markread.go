package task

import (
	"google.golang.org/api/gmail/v1"
)

// MarkRead marks given messages in the Gmail service as read
func MarkRead(srv *gmail.Service, r *gmail.ListMessagesResponse, userId string) error {
	for _, m := range r.Messages {
		modifyMessageRequest := gmail.ModifyMessageRequest{
			RemoveLabelIds: []string{"UNREAD"},
		}
		_, err := srv.Users.Messages.Modify(userId, m.Id, &modifyMessageRequest).Do()
		if err != nil {
			return err
		}
	}

	return nil
}
