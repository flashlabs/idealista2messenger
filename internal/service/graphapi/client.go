package graphapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/flashlabs/idealista2messenger/internal/structs"
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
	PageID          string
}

func NewGraphApiClient(tokenLocation, pageID string) (Client, error) {
	token, err := PageAccessTokenFromFile(tokenLocation)
	if err != nil {
		return Client{}, err
	}

	return Client{PageAccessToken: token, PageID: pageID}, nil
}

func (c Client) SendMessage(recipientID, title, imageURL, link string) (int, error) {
	msg := fmt.Sprintf(template, title, imageURL, link, link)

	return c.sendMessage(recipientID, msg)
}

func (c Client) sendMessage(recipientID string, msg string) (int, error) {
	recipient, err := json.Marshal(structs.Recipient{Id: recipientID})
	if err != nil {
		return 0, fmt.Errorf("error while marshalling json: %w", err)
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, c.apiURL(recipient, msg), nil)
	if err != nil {
		return 0, fmt.Errorf("error while creating http client: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return resp.StatusCode, fmt.Errorf("error while sending a post request: %w", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return resp.StatusCode, fmt.Errorf("error while closing body: %w", err)
	}

	return resp.StatusCode, nil
}

func (c Client) apiURL(recipient []byte, msg string) string {
	return fmt.Sprintf("%s/%s/%s/messages?recipient=%s&messaging_type=RESPONSE&message=%s&access_token=%s",
		apiUrl,
		version,
		c.PageID,
		url.QueryEscape(string(recipient)),
		url.QueryEscape(msg),
		c.PageAccessToken.Token,
	)
}
