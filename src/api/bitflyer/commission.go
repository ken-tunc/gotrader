package bitflyer

const commissionRatePath = "/v1/me/gettradingcommission"

type Commission struct {
	CommissionRate float64 `json:"commission_rate"`
}

type commissionClient struct {
	client *Client
}

func (c *commissionClient) GetCommissionRate(productCode string) (*Commission, error) {
	query := map[string]string{
		"product_code": productCode,
	}

	var commission Commission
	if err := c.client.doRequest(commissionRatePath, "GET", query, nil, &commission); err != nil {
		return nil, err
	}

	return &commission, nil
}
