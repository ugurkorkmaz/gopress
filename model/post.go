package model

type PostCreateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PostCreateResponse struct {
	ID int `json:"id"`
}
type PostDeleteRequest struct {
	ID int `json:"id"`
}
type PostDeleteResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}
type PostUpdateRequest struct {
	ID    int    `json:"id" required:"true"`
	Title string `json:"title" required:"true"`
	Body  string `json:"body" required:"true"`
}

type PostUpdateResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

type PostListResponse struct {
	Posts []Post `json:"data"`
}
type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}
