package jwt

type TokenValidationOptions struct {
	Token     string `json:"token"`
	IsRefresh bool   `json:"is_refresh"`
}
