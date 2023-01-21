package messages

import (
	"github.com/flashlabs/idealista2messenger/internal/process/messages/input"
	"github.com/flashlabs/idealista2messenger/internal/process/messages/task"
)

func Execute(params input.Params) (bool, error) {
	credentials, err := task.Credentials(params.CredentialsFileLocation)
	if err != nil {
		return false, err
	}

	googleClient, err := task.GoogleClient(credentials, params.AccessTokenFileLocation)
	if err != nil {
		return false, err
	}

	srv, err := task.GmailService(googleClient)
	if err != nil {
		return false, err
	}

	messages, err := task.FetchMessages(srv, params.GmailUserId, params.GmailQuery)
	if err != nil {
		return false, err
	}

	err = task.SendMessages(srv, messages, params.PageAccessTokenFileLocation, params.GmailUserId, params.PageId, params.Recipients)
	if err != nil {
		return false, err
	}

	err = task.MarkRead(srv, messages, params.GmailUserId)
	if err != nil {
		return false, err
	}

	return true, nil
}
