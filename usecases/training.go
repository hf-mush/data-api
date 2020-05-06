package usecases

import (
	"time"

	"github.com/shuufujita/data-api/domain/model"
	"github.com/shuufujita/data-api/domain/repository"
)

// TrainingUseCase usecase of training
type TrainingUseCase interface {
	GetLogs(kind string, page int) ([]*model.TrainingLog, error)
	CreateLog(kind string, date string, count int) error
	UpdateLog(objectID string, kind string, date string, count int) error
	DeleteLog(objectID string) error
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

func (tu trainingUseCase) GetLogs(kind string, page int) ([]*model.TrainingLog, error) {
	if kind != "" {
		trainingList, err := tu.repository.GetTrainingLogByKind(page, kind)
		if err != nil {
			return nil, err
		}
		return trainingList, nil
	}

	trainingList, err := tu.repository.GetTrainingLogAll(page)
	if err != nil {
		return nil, err
	}
	return trainingList, nil
}

func (tu trainingUseCase) CreateLog(kind string, date string, count int) error {
	_, err := parseJstWithRFC3339(date)
	if err != nil {
		return err
	}

	err = tu.repository.InsertTrainingLog(kind, date, count)
	if err != nil {
		return err
	}
	return nil
}

func (tu trainingUseCase) UpdateLog(objectID string, kind string, date string, count int) error {
	_, err := parseJstWithRFC3339(date)
	if err != nil {
		return err
	}

	err = tu.repository.UpdateTrainingLog(objectID, kind, date, count)
	if err != nil {
		return err
	}
	return nil
}

func parseJstWithRFC3339(date string) (time.Time, error) {
	return time.Parse(time.RFC3339, date)
}

func (tu trainingUseCase) DeleteLog(objectID string) error {
	err := tu.repository.DeleteTrainingLog(objectID)
	if err != nil {
		return err
	}
	return nil
}
