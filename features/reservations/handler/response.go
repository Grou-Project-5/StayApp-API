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

type HistoryResponse struct {
	RoomName      string `json:"room_name"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	GrossAmount   int64  `json:"gross_amount"`
	OrderID       string `json:"order_id"`
	StatusPayment string `json:"status_payment"`
}

type ListHistoryResponse []HistoryResponse
