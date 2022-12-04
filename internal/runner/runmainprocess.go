package runner

import (
	"fmt"
	"github.com/flashlabs/idealista2messenger/internal/process/mainprocess"
	"github.com/flashlabs/idealista2messenger/internal/process/mainprocess/input"
)

func RunMainProcess() bool {
	fmt.Println("Main process")

	success, err := mainprocess.Execute(input.Input{
		AccessTokenFileLocation:     "config/token.json",
		PageAccessTokenFileLocation: "config/page_access_token.json",
		CredentialsFileLocation:     "config/credentials.json",
		GmailUserId:                 "me",                                             // TODO: move to .env
		GmailQuery:                  "is:unread from:idealista.com",                   // TODO: move to .env
		PageId:                      "302111363180912",                                // TODO: move to .env
		Recipients:                  []string{"8286382568098581", "5612631752117262"}, // TODO: move to .env
	})
	if err != nil {
		fmt.Println(err)

		return false
	}

	return success
}
