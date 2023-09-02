package task

import (
	"errors"
	"fmt"
	"log"

	"google.golang.org/api/gmail/v1"
)

// DisplayMessages reads user messages and prints them to output
func DisplayMessages(srv *gmail.Service, r *gmail.ListMessagesResponse, userId string) error {
	if len(r.Messages) == 0 {
		return errors.New("no messages found")
	}

	fmt.Println("Messages:")
	for _, m := range r.Messages {
		msg, err := srv.Users.Messages.Get(userId, m.Id).Do()
		if err != nil {
			log.Fatalf("Unable to read message details: %v", err)
		}
		fmt.Printf("%s %s: %s\n", msg.Id, msg.LabelIds, msg.Snippet)
	}

	return nil
}
