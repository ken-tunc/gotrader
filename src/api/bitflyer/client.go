package bitflyer

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var baseUrl *url.URL

func init() {
	u, err := url.Parse("https://api.bitflyer.com")
	if err != nil {
		log.Fatalf("cannot parse bitflyer api url: %s", err)
	}
	baseUrl = u
}

type Client struct {
	key        string
	secret     string
	httpClient *http.Client
	wsTimeout  time.Duration

	Balance    *balanceClient
	Commission *commissionClient
	Order      *orderClient
	Realtime   *realtimeClient
}

func NewClient(key, secret string, httpTimeout, wsTimeout time.Duration) *Client {
	client := &Client{
		key:    key,
		secret: secret,
		httpClient: &http.Client{
			Timeout: httpTimeout,
		},
		wsTimeout: wsTimeout,
	}

	client.Balance = &balanceClient{client: client}
	client.Commission = &commissionClient{client: client}
	client.Order = &orderClient{client: client}
	client.Realtime = &realtimeClient{client: client}

	return client
}

func (c *Client) doRequest(path, method string, query map[string]string, body []byte, data interface{}) error {
	uri, err := url.Parse(path)
	if err != nil {
		return err
	}

	endpoint := baseUrl.ResolveReference(uri).String()
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	for key, value := range c.headers(method, req.URL.RequestURI(), body) {
		req.Header.Add(key, value)
	}

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dataBin, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(dataBin) == 0 {
		return nil
	}

	return json.Unmarshal(dataBin, data)
}

func (c *Client) headers(method, endpoint string, body []byte) map[string]string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	msg := timestamp + method + endpoint + string(body)

	mac := hmac.New(sha256.New, []byte(c.secret))
	mac.Write([]byte(msg))

	sign := hex.EncodeToString(mac.Sum(nil))
	return map[string]string{
		"ACCESS-KEY":       c.key,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}
}
