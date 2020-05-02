package usecases

import (
	"github.com/hf-mush/data-api/domain/model"
	"github.com/hf-mush/data-api/domain/repository"
)

// TrainingUseCase usecase of training
type TrainingUseCase interface {
	GetAll() ([]*model.Training, error)
	GetByKind(kind string) ([]*model.Training, error)
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
