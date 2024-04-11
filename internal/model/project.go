package model

import "time"

type Project struct {
	ProjectId int
	UserId    int
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
