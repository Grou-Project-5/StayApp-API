package handler

type FeedbackResponse struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"created_at"`
}

type ListFeedbackResponse []FeedbackResponse