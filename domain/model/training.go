package model

// Training training struct
type Training struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
	Kind  string `json:"kind"`
}
