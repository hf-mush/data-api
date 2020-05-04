package usecases

import (
	"github.com/shuufujita/data-api/domain/model"
	"github.com/shuufujita/data-api/domain/repository"
)

// TrainingUseCase usecase of training
type TrainingUseCase interface {
	GetLogAll() ([]*model.Training, error)
	GetLogByKind(kind string) ([]*model.Training, error)
	GetKindByKindTag(kind string) (*model.TrainingKind, error)
	CreateLog(trainingKindID int64, tag string, count int) error
	UpdateLog(trainingLogID, trainingKindID int64, tag string, count int) error
	DeleteLog(trainingLogID int64) error
}

type trainingUseCase struct {
	repository repository.TrainingRepository
}

// NewTrainingUseCase return training usecase entity
func NewTrainingUseCase(tr repository.TrainingRepository) TrainingUseCase {
	return &trainingUseCase{
		repository: tr,
	}
}

func (tu trainingUseCase) GetLogAll() ([]*model.Training, error) {
	trainingList, err := tu.repository.GetTrainingLogAll()
	if err != nil {
		return nil, err
	}
	return trainingList, nil
}

func (tu trainingUseCase) GetLogByKind(kind string) ([]*model.Training, error) {
	trainingList, err := tu.repository.GetTrainingLogByKind(kind)
	if err != nil {
		return nil, err
	}
	return trainingList, nil
}

func (tu trainingUseCase) GetKindByKindTag(kind string) (*model.TrainingKind, error) {
	trainingKind, err := tu.repository.GetTrainingKindByKindTag(kind)
	if err != nil {
		return nil, err
	}
	return trainingKind, nil
}

func (tu trainingUseCase) CreateLog(trainingKindID int64, tag string, count int) error {
	err := tu.repository.InsertTrainingLog(trainingKindID, tag, count)
	if err != nil {
		return err
	}
	return nil
}

func (tu trainingUseCase) UpdateLog(trainingLogID, trainingKindID int64, tag string, count int) error {
	err := tu.repository.UpdateTrainingLog(trainingLogID, trainingKindID, tag, count)
	if err != nil {
		return err
	}
	return nil
}

func (tu trainingUseCase) DeleteLog(trainingLogID int64) error {
	err := tu.repository.DeleteTrainingLog(trainingLogID)
	if err != nil {
		return err
	}
	return nil
}
