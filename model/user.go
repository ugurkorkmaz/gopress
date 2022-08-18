package model

type UserDeleteResponse struct {
	Message string `json:"message"`
}
type UserDeleteRequest struct {
	ID string `json:"id"`
}

type UserListResponse struct {
	Users []Users `json:"users"`
}
type UserListRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
type Users struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UserShowResponse struct {
	Users Users `json:"users"`
}
type UserShowRequest struct {
	Users Users `json:"users"`
}
type UserUpdateResponse struct {
	Message string `json:"message"`
}
