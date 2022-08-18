package model

type LoginRequest struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}
type LoginResponse struct {
	Token string `json:"token"`
}
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
type RegisterResponse struct {
	Token  string `json:"token"`
	UserID int    `json:"user_id"`
}
