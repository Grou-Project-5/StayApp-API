package handler

type FeedbackResponse struct {
	ID        uint    `json:"id"`
	UserName  string  `json:"user_name"`
	Rating    float64 `json:"rating"`
	Comment   string  `json:"comment"`
	CreatedAt string  `json:"created_at"`
}

type ListFeedbackResponse []FeedbackResponse