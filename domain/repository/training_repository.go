package repository

import "github.com/shuufujita/data-api/domain/model"

// TrainingRepository db.training repository
type TrainingRepository interface {
	GetTrainingAll() ([]*model.Training, error)
	GetTrainingByKind(kind string) ([]*model.Training, error)
	GetTrainingKindByKindTag(tag string) (*model.TrainingKind, error)
	InsertTrainingLog(trainingKindID int64, date string, count int) error
}
