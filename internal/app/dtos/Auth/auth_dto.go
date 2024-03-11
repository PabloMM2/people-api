package authDto

type AuthRequest struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type TokenResponse struct {
	Token *string `json:"token"`
}
