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
	pageId          = "302111363180912"
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

func SendMessages(srv *gmail.Service, r *gmail.ListMessagesResponse) error {
	token, err := pkgToken.PageAccessTokenFromFile("config/page_access_token.json")
	if err != nil {
		return err
	}

	recipients := []string{"8286382568098581", "5612631752117262"}
	user := "me"

	fmt.Printf("Sending %d message(s) to %s\n", len(r.Messages), recipients)
	for _, m := range r.Messages {
		msg, err := srv.Users.Messages.Get(user, m.Id).Format("raw").Do()
		if err != nil {
			log.Fatalf("Unable to read message details: %v", err)
		}

		imageUrl, link := parseMessage(msg.Raw)
		messengerMsg := fmt.Sprintf(messageTemplate, msg.Snippet, imageUrl, link, link)
		for _, recipientId := range recipients {
			rcp := structs.Recipient{Id: recipientId}
			recipient, _ := json.Marshal(rcp)
			resp, err := http.Post(getApiUrl(recipient, messengerMsg, token.Token), "", nil)
			if err != nil {
				log.Fatalf("Unable to send message details: %v", err)
			}
			_ = resp.Body.Close()
			fmt.Printf("Message %s to %s sent\n", msg.Id, rcp.Id)
		}
	}

	return nil
}

func getApiUrl(recipient []byte, messengerMsg, token string) string {
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

	r, _ = regexp.Compile("blur/(.+)/0")
	urlPart := r.FindString(raw)
	if urlPart == "" {
		urlPart = "blur/500_375_mq/0"
	}

	r, _ = regexp.Compile("([a-z0-9]+)/([a-z0-9]+)/([a-z0-9]+)/([a-z0-9]+).jpg")
	imageUrl = r.FindString(raw)

	imageUrlTemplate := "https://img3.idealista.com/%s/id.pro.es.image.master/%s"

	return fmt.Sprintf(imageUrlTemplate, urlPart, imageUrl), link
}
