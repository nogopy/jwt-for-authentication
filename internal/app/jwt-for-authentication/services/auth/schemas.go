package auth

type RegisterBodySchema struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResultSchema struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	ExpiredAt    int    `json:"expires_at"`
	UserId       int    `json:"user_id"`
}
