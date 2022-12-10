package graphapi

import (
	"encoding/json"
	"fmt"
	"github.com/flashlabs/idealista2messenger/internal/structs"
	"net/http"
	"net/url"
)

const (
	apiUrl   = "https://graph.facebook.com"
	version  = "v15.0"
	template = `
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

type Client struct {
	PageAccessToken *PageAccessToken
	PageId          string
}

func NewGraphApiClient(PageAccessTokenFileLocation, PageId string) (Client, error) {
	token, err := PageAccessTokenFromFile(PageAccessTokenFileLocation)
	if err != nil {
		return Client{}, err
	}

	return Client{PageAccessToken: token, PageId: PageId}, nil
}

func (c Client) SendMessage(recipientId, title, imageURL, link string) (int, error) {
	msg := fmt.Sprintf(template, title, imageURL, link, link)

	return c.sendMessage(recipientId, msg)
}

func (c Client) sendMessage(recipientId string, msg string) (int, error) {
	recipient, _ := json.Marshal(structs.Recipient{Id: recipientId})
	resp, err := http.Post(c.apiURL(recipient, msg), "", nil)
	_ = resp.Body.Close()

	return resp.StatusCode, err
}

func (c Client) apiURL(recipient []byte, msg string) string {
	return fmt.Sprintf("%s/%s/%s/messages?recipient=%s&messaging_type=RESPONSE&message=%s&access_token=%s",
		apiUrl,
		version,
		c.PageId,
		url.QueryEscape(string(recipient)),
		url.QueryEscape(msg),
		c.PageAccessToken.Token,
	)
}
