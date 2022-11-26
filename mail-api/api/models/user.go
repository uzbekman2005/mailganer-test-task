package models

type User struct {
	Name          string `json:"name"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	EmailPassword string `json:"email_password"`
}

type SaveUserInRedis struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	EmailPassword string `json:"email_password"`
	AccessToken   string `json:"access_token"`
}