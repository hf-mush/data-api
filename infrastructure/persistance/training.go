package persistance

import (
	"github.com/hf-mush/data-api/domain/model"
	"github.com/hf-mush/data-api/domain/repository"
)

type trainingPersistance struct{}

// NewTrainingPersistance db.training repository persistance
func NewTrainingPersistance() repository.TrainingRepository {
	return &trainingPersistance{}
}

func (tp trainingPersistance) GetTrainingByKind(kind string) ([]*model.Training, error) {
	trainings := []*model.Training{}
	trainging := &model.Training{
		Date:  "",
		Count: 40,
		Kind:  "pushup",
	}
	trainings = append(trainings, trainging)

	return trainings, nil
}
