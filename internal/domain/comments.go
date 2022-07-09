package domain

import "time"

// Comment: Is a data model for comments

type Comment struct {
	Id      int
	Text    string
	Created time.Time
}
