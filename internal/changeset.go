package internal

import "time"

type Changeset struct {
	Id          int
	Title       string
	Description string
	Create      string
	Modify      string
	Delete      string
	Username    string
	Date        time.Time
}
