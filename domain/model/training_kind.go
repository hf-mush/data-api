package model

// TrainingKind training_kind struct
type TrainingKind struct {
	TrainingKindID int64  `json:"training_kind_id"`
	Tag            string `json:"tag"`
	Name           string `json:"name"`
}
