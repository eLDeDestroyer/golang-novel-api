package dto

type ResponseLogin struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
