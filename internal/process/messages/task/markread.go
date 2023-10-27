package task

import (
	"fmt"
	"sync"

	"google.golang.org/api/gmail/v1"
)

type markAsReadParams struct {
	Server               *gmail.Service
	UserID               string
	MessageID            string
	ModifyMessageRequest *gmail.ModifyMessageRequest
}

type markAsReadOperators struct {
	WaitGroup *sync.WaitGroup
	Limiter   chan int
	Counter   int
}

// MarkRead marks given messages in the Gmail service as read.
func MarkRead(srv *gmail.Service, r *gmail.ListMessagesResponse, userId string) error {
	wg := sync.WaitGroup{}
	limiter := make(chan int, 20)

	defer close(limiter)

	counter := 0

	for _, m := range r.Messages {
		wg.Add(1)
		limiter <- 1
		counter++

		modifyMessageRequest := gmail.ModifyMessageRequest{
			RemoveLabelIds: []string{"UNREAD"},
		}

		go modifyMessage(markAsReadParams{
			Server:               srv,
			UserID:               userId,
			MessageID:            m.Id,
			ModifyMessageRequest: &modifyMessageRequest,
		},
			markAsReadOperators{
				WaitGroup: &wg,
				Limiter:   limiter,
				Counter:   counter,
			})
	}

	wg.Wait()

	return nil
}

func modifyMessage(p markAsReadParams, op markAsReadOperators) {
	defer op.WaitGroup.Done()

	_, err := p.Server.Users.Messages.Modify(p.UserID, p.MessageID, p.ModifyMessageRequest).Do()
	if err != nil {
		fmt.Printf("Unable to mark as read message details: %v", err)
	}

	fmt.Printf("%d) Message %s marked as read\n", op.Counter, p.MessageID)

	<-op.Limiter
}
