package task

import (
	"errors"
	"fmt"
	"google.golang.org/api/gmail/v1"
)

func FetchLabels(srv *gmail.Service) error {
	user := "me"
	r, err := srv.Users.Labels.List(user).Do()
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to retrieve labels: %v", err))
	}
	if len(r.Labels) == 0 {
		return errors.New("no labels found")
	}
	fmt.Println("Labels:")
	for _, l := range r.Labels {
		fmt.Printf("- %s\n", l.Name)
	}

	return nil
}
