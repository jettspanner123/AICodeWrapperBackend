package jwt

type TokenClaims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
}
