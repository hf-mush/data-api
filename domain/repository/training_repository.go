package repository

import "github.com/hf-mush/data-api/domain/model"

// TrainingRepository db.training repository
type TrainingRepository interface {
	GetTrainingByKind(kind string) ([]*model.Training, error)
}
