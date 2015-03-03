package tatslack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	Token string
}

func (c *Client) ChannelHistory(channel string) (*Response, error) {
	//generate url
	u := &url.URL{
		Scheme: "https",
		Host:   "slack.com",
		Path:   "/api/channels.history",
	}
	v := url.Values{}
	v.Set("token", c.Token)
	v.Set("channel", channel)
	u.RawQuery = v.Encode()
	fmt.Println("will fetch channel from urls ", u.String())

	//get resp from slack
	resp, err := http.Get(u.String())
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	r := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}

	return r, nil
}

type Response struct {
	OK        bool       `json:"ok"`
	Messages  []*Message `json:"messages"`
	HasMore   bool       `json:"has_more"`
	IsLimited bool       `json:"is_limited"`
}

type Message struct {
	Type   string `json:"type"`
	TS     string `json:"ts"`
	UserID string `json:"user"`
	Text   string `json:"text"`
}
