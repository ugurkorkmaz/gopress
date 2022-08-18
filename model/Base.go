package model

type Exception struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
