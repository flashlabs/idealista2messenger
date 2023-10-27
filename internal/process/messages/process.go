package messages

import (
	"fmt"

	"github.com/flashlabs/idealista2messenger/internal/process/messages/input"
	"github.com/flashlabs/idealista2messenger/internal/process/messages/task"
)

func Execute(params input.Params) (bool, error) {
	credentials, err := task.Credentials(params.CredentialsFileLocation)
	if err != nil {
		return false, fmt.Errorf("error executing credentials task: %w", err)
	}

	googleClient, err := task.GoogleClient(credentials, params.AccessTokenFileLocation)
	if err != nil {
		return false, fmt.Errorf("error executing google client task: %w", err)
	}

	srv, err := task.GmailService(googleClient)
	if err != nil {
		return false, fmt.Errorf("error executing gmail service task: %w", err)
	}

	messages, err := task.FetchMessages(srv, params.GmailUserId, params.GmailQuery)
	if err != nil {
		return false, fmt.Errorf("error executing fetch messages task: %w", err)
	}

	err = task.SendMessages(srv, messages, params.PageAccessTokenFileLocation, params.GmailUserId, params.PageId, params.Recipients)
	if err != nil {
		return false, fmt.Errorf("error executing send messages task: %w", err)
	}

	err = task.MarkRead(srv, messages, params.GmailUserId)
	if err != nil {
		return false, fmt.Errorf("error executing mark read task: %w", err)
	}

	return true, nil
}
