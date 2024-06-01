package userhdl

type BodyLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ResponseLogin(token string) string {
	return token
}
