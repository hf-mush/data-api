package repository

import "github.com/shuufujita/data-api/domain/model"

// TrainingRepository db.training repository
type TrainingRepository interface {
	GetTrainingLogAll() ([]*model.TrainingLog, error)
	GetTrainingLogByKind(kind string) ([]*model.TrainingLog, error)
	GetTrainingKindByKindTag(tag string) (*model.TrainingKind, error)
	InsertTrainingLog(trainingKindID int64, date string, count int) error
	UpdateTrainingLog(trainingLogID, trainingKindID int64, date string, count int) error
	DeleteTrainingLog(trainingLogID int64) error
}
