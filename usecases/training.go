package usecases

import (
	"github.com/shuufujita/data-api/domain/model"
	"github.com/shuufujita/data-api/domain/repository"
)

// TrainingUseCase usecase of training
type TrainingUseCase interface {
	GetAll() ([]*model.Training, error)
	GetByKind(kind string) ([]*model.Training, error)
	GetKindByKindTag(kind string) (*model.TrainingKind, error)
	CreateLog(trainingKindID int64, tag string, count int) error
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

func (tu trainingUseCase) GetAll() ([]*model.Training, error) {
	trainingList, err := tu.repository.GetTrainingAll()
	if err != nil {
		return nil, err
	}
	return trainingList, nil
}

func (tu trainingUseCase) GetByKind(kind string) ([]*model.Training, error) {
	trainingList, err := tu.repository.GetTrainingByKind(kind)
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
