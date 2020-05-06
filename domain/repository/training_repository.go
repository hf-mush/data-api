package repository

import "github.com/shuufujita/data-api/domain/model"

// TrainingRepository db.training repository
type TrainingRepository interface {
	GetTrainingLogAll() ([]*model.TrainingLog, error)
	GetTrainingLogByKind(kind string) ([]*model.TrainingLog, error)
	InsertTrainingLog(kind string, date string, count int) error
	UpdateTrainingLog(objectID string, kind string, date string, count int) error
	DeleteTrainingLog(objectID string) error
}
