package persistance

import (
	"fmt"

	"github.com/hf-mush/data-api/domain/model"
	"github.com/hf-mush/data-api/domain/repository"
)

type trainingPersistance struct{}

// NewTrainingPersistance db.training repository persistance
func NewTrainingPersistance() repository.TrainingRepository {
	return &trainingPersistance{}
}

func (tp trainingPersistance) GetTrainingByKind(kind string) ([]*model.Training, error) {
	conn := GetConn()
	stmt, err := conn.Prepare(fmt.Sprintf("SELECT tl.date, tk.tag, tl.count FROM training_logs AS tl INNER JOIN training_kinds AS tk ON tl.training_kind_id = tk.training_kind_id WHERE tk.tag = ?"))
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(kind)
	if err != nil {
		return nil, err
	}

	trainings := []*model.Training{}

	var date string
	var tag string
	var count int

	for rows.Next() {
		err := rows.Scan(&date, &tag, &count)
		if err != nil {
			panic(err)
		}
		trainging := &model.Training{
			Date:  date,
			Count: count,
			Kind:  tag,
		}
		trainings = append(trainings, trainging)
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return trainings, nil
}
