package entities

import "time"

type Books struct {
	Id          uint
	Name        string
	Category_id Categories
	Stock       int
	Description string
	Created_at  time.Time
	Updated_at  time.Time
}