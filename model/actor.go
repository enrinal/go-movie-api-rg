package model

type Actor struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Age         int    `json:"age"`
}
