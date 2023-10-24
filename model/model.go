package model

type Pet struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Breed string `json:"breed"`
}
