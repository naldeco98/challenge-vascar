package domain

import "time"

type ReportPost struct {
	Id      string    `json:"id"`
	Message string    `json:"message" binding:"required"`
	UserId  int       `json:"user_id" binding:"required"`
	Created time.Time `json:"created" binding:"required"`
	PostId int       `json:"message_id" binding:"required"`
}

type ReportMessage struct {
	Id        string    `json:"id"`
	Message   string    `json:"message" binding:"required"`
	UserId    int       `json:"user_id" binding:"required"`
	Created   time.Time `json:"created" binding:"required"`
	MessageId int       `json:"message_id" binding:"required"`
}
