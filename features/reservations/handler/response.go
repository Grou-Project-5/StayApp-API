package handler

type Availability struct {
	Availability bool `json:"availability"`
}

type PaymentResponse struct {
	OrderID     string `json:"order_id"`
	RedirectURL string `json:"redirect_url"`
}

type PayStatusResponse struct {
	OrderID       string `json:"order_id"`
	StatusPayment string `json:"payment_status"`
}
