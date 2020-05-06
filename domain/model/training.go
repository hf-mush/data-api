package model

// TrainingLog training struct
type TrainingLog struct {
	ID    string `json:"id"`
	Date  string `json:"date"`
	Count int    `json:"count"`
	Kind  string `json:"kind"`
}
