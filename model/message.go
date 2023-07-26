package model

type MessageRequest struct {
	Channel   string `json:"channel" validate:"required"`
	Recipient string `json:"recipient" validate:"min=2"`
	Content   string `json:"content"`
}
