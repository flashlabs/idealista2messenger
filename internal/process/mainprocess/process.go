package mainprocess

import (
	"github.com/flashlabs/idealista2messenger/internal/process/mainprocess/input"
	"github.com/flashlabs/idealista2messenger/internal/process/mainprocess/task"
)

func Execute(input input.Input) (bool, error) {
	credentials, err := task.LoadCredentials()
	if err != nil {
		return false, err
	}

	googleClient, err := task.GetGoogleClient(credentials, input.AccessTokenFileLocation)
	if err != nil {
		return false, err
	}

	srv, err := task.GetGmailService(googleClient)
	if err != nil {
		return false, err
	}

	messages, err := task.FetchMessages(srv, input.GmailUserId, input.GmailQuery)
	if err != nil {
		return false, err
	}

	//err = task.DisplayMessages(srv, messages, input.GmailUserId)
	//if err != nil {
	//	return false, err
	//}

	err = task.SendMessages(srv, messages, input.PageAccessTokenFileLocation, input.GmailUserId, input.PageId, input.Recipients)
	if err != nil {
		return false, err
	}

	err = task.MarkRead(srv, messages, input.GmailUserId)
	if err != nil {
		return false, err
	}

	return true, nil
}
