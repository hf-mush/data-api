package persistance

import (
	"database/sql"
	"fmt"

	"github.com/shuufujita/data-api/domain/model"
	"github.com/shuufujita/data-api/domain/repository"
)

type trainingPersistance struct{}

// NewTrainingPersistance db.training repository persistance
func NewTrainingPersistance() repository.TrainingRepository {
	return &trainingPersistance{}
}

func (tp trainingPersistance) GetTrainingLogAll() ([]*model.Training, error) {
	conn := GetConn()
	stmt, err := conn.Prepare(fmt.Sprintf("SELECT tl.training_log_id, tl.date, tk.tag, tl.count FROM training_logs AS tl INNER JOIN training_kinds AS tk ON tl.training_kind_id = tk.training_kind_id"))
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	return aggregateTrainingLogs(rows)
}

func (tp trainingPersistance) GetTrainingLogByKind(kind string) ([]*model.Training, error) {
	conn := GetConn()
	stmt, err := conn.Prepare(fmt.Sprintf("SELECT tl.training_log_id, tl.date, tk.tag, tl.count FROM training_logs AS tl INNER JOIN training_kinds AS tk ON tl.training_kind_id = tk.training_kind_id WHERE tk.tag = ?"))
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(kind)
	if err != nil {
		return nil, err
	}

	return aggregateTrainingLogs(rows)
}

func aggregateTrainingLogs(rows *sql.Rows) ([]*model.Training, error) {
	trainings := []*model.Training{}

	var trainingLogID int64
	var date string
	var tag string
	var count int

	for rows.Next() {
		err := rows.Scan(&trainingLogID, &date, &tag, &count)
		if err != nil {
			panic(err)
		}
		trainging := &model.Training{
			ID:    trainingLogID,
			Date:  date,
			Count: count,
			Kind:  tag,
		}
		trainings = append(trainings, trainging)
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return trainings, nil
}

func (tp trainingPersistance) GetTrainingKindByKindTag(kind string) (*model.TrainingKind, error) {
	conn := GetConn()

	stmt, err := conn.Prepare("SELECT * FROM training_kinds WHERE tag = ?")
	if err != nil {
		return nil, err
	}

	var trainingKindID int64
	var tag string
	var name string
	if err := stmt.QueryRow(kind).Scan(&trainingKindID, &tag, &name); err != nil {
		return nil, err
	}

	return &model.TrainingKind{
		TrainingKindID: trainingKindID,
		Tag:            tag,
		Name:           name,
	}, nil
}

func (tp trainingPersistance) InsertTrainingLog(trainingKindID int64, date string, count int) error {
	conn := GetConn()
	stmt, err := conn.Prepare("INSERT INTO training_logs (training_kind_id, date, count) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(trainingKindID, date, count)
	if err != nil {
		return err
	}

	return nil
}

func (tp trainingPersistance) UpdateTrainingLog(trainingLogID, trainingKindID int64, date string, count int) error {
	conn := GetConn()
	stmt, err := conn.Prepare("UPDATE training_logs SET training_kind_id = ?, date = ?, count = ? WHERE training_log_id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(trainingKindID, date, count, trainingLogID)
	if err != nil {
		return err
	}

	return nil
}

func (tp trainingPersistance) DeleteTrainingLog(trainingLogID int64) error {
	conn := GetConn()
	stmt, err := conn.Prepare("DELETE FROM training_logs WHERE training_log_id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(trainingLogID)
	if err != nil {
		return err
	}

	return nil
}
