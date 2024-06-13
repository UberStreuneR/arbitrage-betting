package models

type UserLoginSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthConfig struct {
	Secret string `json:"secret"`
}
