package models


type SendNewsToSupscribersReq struct {
	To   []*Subscriber `json:"subscribers"`
	News string        `json:"news"`
}

type Subscriber struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type StatusInfo struct {
	Message string `json:"message"`
}

type SendEmailConfig struct {
	Email    string
	Passwrod string
}

type SendScheduledEmailsReq struct {
	News           string        `json:"news"`
	To             []*Subscriber `json:"subscribers"`
	MinutsAfter    int           `json:"minuts_after"`
	SenderEmail    string        `json:"sender_email"`
	EmailPaassword string        `json:"email_password"`
}

type SendScheduledEmailsApiReq struct {
	News        string        `json:"news"`
	To          []*Subscriber `json:"subscribers"`
	MinutsAfter int
}

type GetScheduledMessagesRes struct {
	News           string      `json:"news"`
	To             *Subscriber `json:"subscribers"`
	SenderEmail    string      `json:"sender_email"`
	EmailPaassword string      `json:"email_password"`
}
