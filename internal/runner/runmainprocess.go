package runner

import (
	"fmt"
	"github.com/flashlabs/idealista2messenger/internal/initializer"
	"github.com/flashlabs/idealista2messenger/internal/process/mainprocess"
	"github.com/flashlabs/idealista2messenger/internal/process/mainprocess/input"
	"os"
	"strings"
)

func RunMainProcess(config *initializer.Config) bool {
	fmt.Println("Main process")

	success, err := mainprocess.Execute(input.Input{
		AccessTokenFileLocation:     config.Google.AccessTokenFile,
		CredentialsFileLocation:     config.Google.CredentialsFile,
		PageAccessTokenFileLocation: config.Messenger.PageAccessTokenFile,
		GmailUserId:                 os.Getenv("GMAIL_USER_ID"),
		GmailQuery:                  os.Getenv("GMAIL_QUERY"),
		PageId:                      os.Getenv("FB_PAGE_ID"),
		Recipients:                  strings.Split(os.Getenv("FB_PAGE_RECIPIENTS"), ","),
	})
	if err != nil {
		fmt.Println(err)

		return false
	}

	return success
}
