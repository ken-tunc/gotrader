package bitflyer

import "encoding/json"

const (
	sendOrderPath = "/v1/me/sendchildorder"
	listOrderPath = "/v1/me/getchildorders"
)

type OrderRequest struct {
	ProductCode    string  `json:"product_code"`
	ChildOrderType string  `json:"child_order_type"`
	Side           string  `json:"side"`
	Price          int     `json:"price"`
	Size           float64 `json:"size"`
	MinuteToExpire int     `json:"minute_to_expire"`
	TimeInForce    string  `json:"time_in_force"`
}

type SendOrderResponse struct {
	ChildOrderAcceptanceID string `json:"child_order_acceptance_id"`
}

type Order struct {
	ID                     int     `json:"id"`
	ChildOrderID           string  `json:"child_order_id"`
	ProductCode            string  `json:"product_code"`
	Side                   string  `json:"side"`
	ChildOrderType         string  `json:"child_order_type"`
	Price                  int     `json:"price"`
	AveragePrice           int     `json:"average_price"`
	Size                   float64 `json:"size"`
	ChildOrderState        string  `json:"child_order_state"`
	ExpireDate             string  `json:"expire_date"`
	ChildOrderDate         string  `json:"child_order_date"`
	ChildOrderAcceptanceID string  `json:"child_order_acceptance_id"`
	OutstandingSize        int     `json:"outstanding_size"`
	CancelSize             int     `json:"cancel_size"`
	ExecutedSize           float64 `json:"executed_size"`
	TotalCommission        int     `json:"total_commission"`
}

type orderClient struct {
	client *Client
}

func (o *orderClient) SendOrder(order *OrderRequest) (*SendOrderResponse, error) {
	body, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	var data *SendOrderResponse
	if err = o.client.doRequest(sendOrderPath, "POST", map[string]string{}, body, data); err != nil {
		return nil, err
	}

	return data, nil
}

func (o *orderClient) ListOrders(productCode, childOrderAcceptanceId string) ([]Order, error) {
	query := map[string]string{
		"product_code": productCode,
	}
	if childOrderAcceptanceId != "" {
		query["child_order_acceptance_id"] = childOrderAcceptanceId
	}

	var orders []Order
	if err := o.client.doRequest(listOrderPath, "GET", query, nil, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}
