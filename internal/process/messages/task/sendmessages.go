package task

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime/quotedprintable"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"google.golang.org/api/gmail/v1"

	"github.com/flashlabs/idealista2messenger/internal/service/graphapi"
)

type sendMessageParams struct {
	Message     *gmail.Message
	Client      graphapi.Client
	Server      *gmail.Service
	GmailUserID string
	Recipients  []string
}

type sendMessageOperators struct {
	WaitGroup *sync.WaitGroup
	Limiter   chan int
	Counter   int
}

func SendMessages(srv *gmail.Service, r *gmail.ListMessagesResponse, pageAccessToken, gmailUserId, pageId string, recipients []string) error {
	fmt.Printf("Sending %d message(s) to %s\n", len(r.Messages), recipients)

	c, err := graphapi.NewGraphApiClient(pageAccessToken, pageId)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	limiter := make(chan int, 20)
	defer close(limiter)

	counter := 0
	for _, m := range r.Messages {
		wg.Add(1)
		limiter <- 1
		counter++

		go sendMessage(sendMessageParams{
			Message:     m,
			Client:      c,
			Server:      srv,
			GmailUserID: gmailUserId,
			Recipients:  recipients,
		}, sendMessageOperators{
			WaitGroup: &wg,
			Limiter:   limiter,
			Counter:   counter,
		})
	}

	wg.Wait()

	return nil
}

func sendMessage(p sendMessageParams, op sendMessageOperators) {
	defer op.WaitGroup.Done()

	msg, err := p.Server.Users.Messages.Get(p.GmailUserID, p.Message.Id).Format("raw").Do()
	if err != nil {
		log.Fatalf("Unable to read message details: %v", err)
	}

	imageUrl, link := parseMessage(msg.Raw)
	for _, recipientId := range p.Recipients {
		status, err := p.Client.SendMessage(recipientId, msg.Snippet, imageUrl, link)

		if err != nil {
			fmt.Printf("Unable to send message details: %v", err)
			continue
		}

		if status != http.StatusOK {
			fmt.Printf("Message %s to %s not sent, status %d\n", msg.Id, recipientId, status)
			continue
		}

		fmt.Printf("%d) Message %s to %s sent\n", op.Counter, msg.Id, recipientId)
	}

	<-op.Limiter
}

func parseMessage(rawMessage string) (imageUrl, link string) {
	raw := decodePayload(rawMessage)

	r, _ := regexp.Compile("https://www.idealista.com/en/inmueble/([0-9]+)/")
	link = r.FindString(raw)

	r, _ = regexp.Compile("blur/([a-zA-Z0-9_]+)/0")
	urlPart := r.FindString(raw)
	if urlPart == "" {
		urlPart = "blur/500_375_mq/0"
	}

	// the less in pattern, the better
	r, _ = regexp.Compile("([a-z0-9]+)/([a-z0-9]+)/([a-z0-9]+)/([a-z0-9]+).j")
	imageUrl = r.FindString(raw)

	imageUrlTemplate := "https://img3.idealista.com/%s/id.pro.es.image.master/%spg"

	return fmt.Sprintf(imageUrlTemplate, urlPart, imageUrl), link
}

func decodePayload(rawMessage string) string {
	decoded, _ := base64.RawURLEncoding.DecodeString(rawMessage)
	raw := string(decoded)

	// remove the headers parts causing issues
	raw = raw[strings.Index(raw, "<!doctype"):]

	return decodeQuotedPrintable(raw)
}

func decodeQuotedPrintable(raw string) string {
	stringsReader := strings.NewReader(raw)
	qpReader := quotedprintable.NewReader(stringsReader)
	b, _ := io.ReadAll(qpReader)

	return string(b)
}
