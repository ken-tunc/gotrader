package bitflyer

import (
	"net/http"
	"time"
)

type Client struct {
	key        string
	secret     string
	httpClient *http.Client
	wsTimeout  time.Duration

	Realtime *RealtimeClient
}

func NewClient(key string, secret string, httpTimeout time.Duration, wsTimeout time.Duration) *Client {
	client := &Client{
		key:    key,
		secret: secret,
		httpClient: &http.Client{
			Timeout: httpTimeout,
		},
		wsTimeout: wsTimeout,
	}

	client.Realtime = &RealtimeClient{client: client}

	return client
}
