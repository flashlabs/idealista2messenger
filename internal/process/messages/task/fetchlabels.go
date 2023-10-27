package task

import (
	"errors"
	"fmt"

	"google.golang.org/api/gmail/v1"
)

var (
	ErrNoLabelsFound = errors.New("no labels found")
)

// FetchLabels fetches user labels and prints them to output.
func FetchLabels(srv *gmail.Service) error {
	user := "me"
	r, err := srv.Users.Labels.List(user).Do()

	if err != nil {
		return fmt.Errorf("unable to retrieve labels: %w", err)
	}

	if len(r.Labels) == 0 {
		return ErrNoLabelsFound
	}

	fmt.Println("Labels:")

	for _, l := range r.Labels {
		fmt.Printf("- %s\n", l.Name)
	}

	return nil
}
