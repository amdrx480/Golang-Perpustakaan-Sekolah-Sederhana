package models

//dari pada membuat satu2 responnya lebih baik membuat generic
type Response[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}
