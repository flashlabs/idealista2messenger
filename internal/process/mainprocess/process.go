package mainprocess

import (
	"github.com/flashlabs/idealista2messenger/internal/process/mainprocess/task"
)

func Execute() (bool, error) {
	credentials, err := task.LoadCredentials()
	if err != nil {
		return false, err
	}

	client, err := task.GetClient(credentials)
	if err != nil {
		return false, err
	}

	srv, err := task.GetGmailService(client)
	if err != nil {
		return false, err
	}

	messages, err := task.FetchMessages(srv)
	if err != nil {
		return false, err
	}

	err = task.DisplayMessages(srv, messages)
	if err != nil {
		return false, err
	}

	err = task.MarkRead(srv, messages)
	if err != nil {
		return false, err
	}

	return true, nil
}
