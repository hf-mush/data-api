package usecases

import (
	"time"

	"github.com/shuufujita/data-api/domain/model"
	"github.com/shuufujita/data-api/domain/repository"
)

// TrainingUseCase usecase of training
type TrainingUseCase interface {
	GetLogs(kind string) ([]*model.TrainingLog, error)
	CreateLog(kind string, date string, count int) error
	UpdateLog(trainingLogID int64, kind string, date string, count int) error
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

func (tu trainingUseCase) GetLogs(kind string) ([]*model.TrainingLog, error) {
	if kind != "" {
		trainingList, err := tu.repository.GetTrainingLogByKind(kind)
		if err != nil {
			return nil, err
		}
		return trainingList, nil
	}

	trainingList, err := tu.repository.GetTrainingLogAll()
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

	trainingKind, err := tu.repository.GetTrainingKindByKindTag(kind)
	if err != nil {
		return err
	}

	err = tu.repository.InsertTrainingLog(trainingKind.TrainingKindID, date, count)
	if err != nil {
		return err
	}
	return nil
}

func (tu trainingUseCase) UpdateLog(trainingLogID int64, kind string, date string, count int) error {
	_, err := parseJstWithRFC3339(date)
	if err != nil {
		return err
	}

	trainingKind, err := tu.repository.GetTrainingKindByKindTag(kind)
	if err != nil {
		return err
	}

	err = tu.repository.UpdateTrainingLog(trainingLogID, trainingKind.TrainingKindID, date, count)
	if err != nil {
		return err
	}
	return nil
}

func parseJstWithRFC3339(date string) (time.Time, error) {
	return time.Parse(time.RFC3339, date)
}

func (tu trainingUseCase) DeleteLog(trainingLogID int64) error {
	err := tu.repository.DeleteTrainingLog(trainingLogID)
	if err != nil {
		return err
	}
	return nil
}
