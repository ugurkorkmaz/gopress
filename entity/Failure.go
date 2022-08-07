package entity

type Failure struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   error  `json:"error"`
}
