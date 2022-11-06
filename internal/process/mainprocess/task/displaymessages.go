package task

import (
	"errors"
	"fmt"
	"google.golang.org/api/gmail/v1"
	"log"
)

func DisplayMessages(srv *gmail.Service, r *gmail.ListMessagesResponse) error {
	if len(r.Messages) == 0 {
		return errors.New("no messages found")
	}

	fmt.Println("Messages:")
	user := "me"
	for _, m := range r.Messages {
		msg, err := srv.Users.Messages.Get(user, m.Id).Do()
		if err != nil {
			log.Fatalf("Unable to read message details: %v", err)
		}
		fmt.Printf("%s %s: %s\n", msg.Id, msg.LabelIds, msg.Snippet)

		//modifyMessageRequest := gmail.ModifyMessageRequest{
		//	RemoveLabelIds: []string{"UNREAD"},
		//}
		//_, err = srv.Users.Messages.Modify(user, msg.Id, &modifyMessageRequest).Do()
		//if err != nil {
		//	return err
		//}
	}

	return nil
}
