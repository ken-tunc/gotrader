package bitflyer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ken-tunc/gotrader/src/api"
)

const (
	apiSchema         = "wss"
	apiHost           = "ws.lightstream.bitflyer.com"
	apiPath           = "/json-rpc"
	apiJsonRPCVersion = "2.0"

	RFC3339local = "2006-01-02T15:04:05Z"
)

type SubscribeParams struct {
	Channel string `json:"channel"`
}

type Ticker struct {
	ProductCode     string  `json:"product_code"`
	Timestamp       string  `json:"timestamp"`
	TickID          int     `json:"tick_id"`
	BestBid         float64 `json:"best_bid"`
	BestAsk         float64 `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     float64 `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   float64 `json:"total_ask_depth"`
	Ltp             float64 `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

func (t *Ticker) MidPrice() float64 {
	return (t.BestBid + t.BestAsk) / 2
}

func (t *Ticker) DateTime() time.Time {
	dateTime, err := time.ParseInLocation(RFC3339local, t.Timestamp, time.Local)
	if err != nil {
		log.Printf("cannot parse ticker timestamp: got=%s", t.Timestamp)
	}
	return dateTime
}

type realtimeClient struct {
	client *Client
	wsConn *websocket.Conn
}

func (r *realtimeClient) SubscribeTicker(productCode string, ch chan<- Ticker) {
	u := url.URL{Scheme: apiSchema, Host: apiHost, Path: apiPath}

RETRY:
	for {
		if r.wsConn == nil {
			conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				log.Printf("cannot connect to %s: %s", u.String(), err)
				continue RETRY
			}
			r.wsConn = conn

			channel := fmt.Sprintf("lightning_ticker_%s", productCode)
			msg := &api.JsonRPC2{
				Version: apiJsonRPCVersion,
				Method:  "subscribe",
				Params:  &SubscribeParams{channel},
			}

			_ = r.wsConn.SetWriteDeadline(time.Now().Add(r.client.wsTimeout))
			if err = r.wsConn.WriteJSON(msg); err != nil {
				log.Printf("cannot send message to %s: message=%v, error=%s", u.String(), msg, err)
				_ = r.wsConn.Close()
				r.wsConn = nil
				continue RETRY
			}
		}

		for {
			received := new(api.JsonRPC2)

			_ = r.wsConn.SetReadDeadline(time.Now().Add(r.client.wsTimeout))
			if err := r.wsConn.ReadJSON(received); err != nil {
				log.Printf("cannot recieve message from %s: %s", u.String(), err)
				_ = r.wsConn.Close()
				r.wsConn = nil
				continue RETRY
			}

			if received.Method == "channelMessage" {
				switch v := received.Params.(type) {
				case map[string]interface{}:
					for key, bin := range v {
						if key == "message" {
							tickerBin, err := json.Marshal(bin)
							if err != nil {
								log.Println(err)
								continue RETRY
							}

							var ticker Ticker
							if err = json.Unmarshal(tickerBin, &ticker); err != nil {
								log.Println(err)
								continue RETRY
							}
							ch <- ticker
						}
					}
				}
			}
		}
	}
}
