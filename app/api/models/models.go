package models

import "time"

type SendEmailWithSupscribersReq struct {
	To   []*Subscriber `json:"subscribers"`
	News string        `json:"news"`
}

type Subscriber struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Birthday  time.Time `json:"birthday"`
	Email     string    `json:"email"`
}

type StatusInfo struct {
	Message string `json:"message"`
}
