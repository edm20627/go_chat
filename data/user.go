package data

import "time"

type Session struct {
	Id       id
	Uuid     string
	Email    string
	UserId   int
	CreateAt time.Time
}
