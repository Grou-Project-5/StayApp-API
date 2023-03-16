package handler

type Availability struct {
	Availability bool `json:"availability"`
}

type PaymentResponse struct {
	OrderID     string `json:"order_id"`
	RedirectURL string `json:"redirect_url"`
}
