package task

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/flashlabs/idealista2messenger/internal/structs"
	pkgToken "github.com/flashlabs/idealista2messenger/internal/token"
	"google.golang.org/api/gmail/v1"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

const (
	graphApiUrl     = "https://graph.facebook.com"
	apiVersion      = "v15.0"
	messageTemplate = `
{
    "attachment":{
      "type":"template",
      "payload":{
        "template_type":"generic",
        "elements":[
           {
            "title":"%s",
            "image_url":"%s",
            "subtitle":"Delivered with love by flashlabs ♥",
            "default_action": {
              "type": "web_url",
              "url": "%s",
              "webview_height_ratio": "tall"
            },
            "buttons":[
              {
                "type":"web_url",
                "url":"%s",
                "title":"View Offer"
              }             
            ]      
          }
        ]
      }
    }
  }
`
)

func SendMessages(srv *gmail.Service, r *gmail.ListMessagesResponse, pageAccessToken, gmailUserId, pageId string, recipients []string) error {
	token, err := pkgToken.PageAccessTokenFromFile(pageAccessToken)
	if err != nil {
		return err
	}

	fmt.Printf("Sending %d message(s) to %s\n", len(r.Messages), recipients)
	for _, m := range r.Messages {
		msg, err := srv.Users.Messages.Get(gmailUserId, m.Id).Format("raw").Do()
		if err != nil {
			log.Fatalf("Unable to read message details: %v", err)
		}

		imageUrl, link := parseMessage(msg.Raw)
		messengerMsg := fmt.Sprintf(messageTemplate, msg.Snippet, imageUrl, link, link)
		for _, recipientId := range recipients {
			err = sendMessage(recipientId, messengerMsg, token, pageId)
			if err != nil {
				log.Fatalf("Unable to send message details: %v", err)
				continue
			}

			fmt.Printf("Message %s to %s sent\n", msg.Id, recipientId)
		}
	}

	return nil
}

func sendMessage(recipientId string, messengerMsg string, token *structs.PageAccessToken, pageId string) error {
	recipient, _ := json.Marshal(structs.Recipient{Id: recipientId})
	resp, err := http.Post(getApiUrl(recipient, messengerMsg, token.Token, pageId), "", nil)
	fmt.Println(resp.Status)
	_ = resp.Body.Close()
	if err != nil {
		return err
	}

	return nil
}

func getApiUrl(recipient []byte, messengerMsg, token, pageId string) string {
	apiUrl := fmt.Sprintf("%s/%s/%s/messages?recipient=%s&messaging_type=RESPONSE&message=%s&access_token=%s",
		graphApiUrl,
		apiVersion,
		pageId,
		url.QueryEscape(string(recipient)),
		url.QueryEscape(messengerMsg),
		token,
	)
	return apiUrl
}

func parseMessage(rawMessage string) (imageUrl, link string) {
	decoded, _ := base64.RawURLEncoding.DecodeString(rawMessage)
	raw := string(decoded)

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
