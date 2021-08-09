package bitflyer

import (
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	key := "apiKey"
	secret := "apiSecret"
	timeout := time.Second
	client := NewClient(key, secret, timeout, timeout)

	if client.key != key {
		t.Errorf("unexpected key: expected=%s, actual=%s", key, client.key)
	}
	if client.secret != secret {
		t.Errorf("unexpected secret: expected=%s, actual=%s", secret, client.secret)
	}
	if client.wsTimeout != timeout {
		t.Errorf("unexpected wsTimeout: expected=%s, actual=%s", timeout, client.wsTimeout)
	}
	if client.Realtime == nil {
		t.Error("Realtime client is not initialized.")
	}
}
