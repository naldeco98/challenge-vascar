package domain

import "time"

type ReportPost struct {
	Id      int
	Reason  string
	UserId  int
	Created time.Time
	PostId  int
}

type ReportComment struct {
	Id        int
	Reason    string
	UserId    int
	Created   time.Time
	CommentId int
}
