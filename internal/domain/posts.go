package domain

import "time"

// Post: Is a data model for posts

type Post struct {
	Id      int
	Text    string
	Created time.Time
}
