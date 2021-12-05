package models

type Tweet struct {
	Base
	UserID  uint   `json:"userId"`
	Content string `json:"content"`
}

type ReceiveTweet struct {
	Content string `validate:"required,max=140"`
}

type ResponseTweet struct {
	Tweets *[]Tweet `json:"tweets"`
}
