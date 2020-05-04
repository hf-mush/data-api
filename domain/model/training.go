package model

// Training training struct
type Training struct {
	ID    int64  `json:"id"`
	Date  string `json:"date"`
	Count int    `json:"count"`
	Kind  string `json:"kind"`
}
