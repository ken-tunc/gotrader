package bitflyer

const balancePath = "/v1/me/getbalance"

type Balance struct {
	CurrentCode string  `json:"currency_code"`
	Amount      float64 `json:"amount"`
	Available   float64 `json:"available"`
}

type balanceClient struct {
	client *Client
}

func (b *balanceClient) GetBalances() ([]Balance, error) {
	var balances []Balance
	if err := b.client.doRequest(balancePath, "GET", map[string]string{}, nil, &balances); err != nil {
		return nil, err
	}
	return balances, nil
}
