package bitflyer

import (
	"testing"
	"time"
)

func TestTicker_DateTime(t *testing.T) {
	ticker := &Ticker{Timestamp: "2021-01-01T12:34:56.0000000Z"}
	expected := time.Date(2021, time.January, 1, 12, 34, 56, 0, time.Local)
	if ticker.DateTime() != expected {
		t.Errorf("unexpected DateTime(): expected=%s, actual=%s", expected, ticker.DateTime())
	}
}

func TestTicker_GetMidPrice(t *testing.T) {
	ticker := &Ticker{BestAsk: 1000, BestBid: 2000}
	var expected float64 = 1500
	if ticker.MidPrice() != expected {
		t.Errorf("unexpected MidPrice(): expected=%f, actual=%f", expected, ticker.MidPrice())
	}
}

func TestRealtimeClient_SubscribeTicker(t *testing.T) {
	client := NewClient("", "", time.Second, 3*time.Second)

	tests := []struct {
		name        string
		productCode string
	}{
		{
			name:        "subscribe BTC_JPY",
			productCode: "BTC_JPY",
		},
		{
			name:        "subscribe ETH_JPY",
			productCode: "ETH_JPY",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeout := time.After(10 * time.Second)
			done := make(chan bool)

			go func() {
				ch := make(chan Ticker, 1)
				go client.Realtime.SubscribeTicker(tt.productCode, ch)
				for i := 0; i < 3; i++ {
					t.Log(<-ch)
				}

				done <- true
			}()

			select {
			case <-timeout:
				t.Fatal("test timed out")
			case <-done:
			}
		})
	}
}
