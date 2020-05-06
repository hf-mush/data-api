package persistance

import (
	"os"

	"github.com/shuufujita/data-api/domain/model"
	"github.com/shuufujita/data-api/domain/repository"
	"gopkg.in/mgo.v2/bson"
)

type trainingPersistance struct{}

// NewTrainingPersistance db.training repository persistance
func NewTrainingPersistance() repository.TrainingRepository {
	return &trainingPersistance{}
}

var offsetSize = 50

func (tp trainingPersistance) GetTrainingLogAll(page int) ([]*model.TrainingLog, error) {
	sess, err := ConnectMongoDB()
	if err != nil {
		return nil, err
	}
	defer CloseMongoDB(sess)

	col := sess.DB(os.Getenv("MONGO_DBNAME")).C("training_logs")
	query := bson.M{}
	logs := []*model.MgoTrainingLog{}

	skip := getSkipCount(page)
	err = col.Find(query).Skip(skip).Limit(offsetSize).All(&logs)
	if err != nil {
		return nil, err
	}

	return aggregateTrainingLogs(logs)
}

func (tp trainingPersistance) GetTrainingLogByKind(page int, kind string) ([]*model.TrainingLog, error) {
	sess, err := ConnectMongoDB()
	if err != nil {
		return nil, err
	}
	defer CloseMongoDB(sess)

	col := sess.DB(os.Getenv("MONGO_DBNAME")).C("training_logs")
	query := bson.M{
		"kind": kind,
	}
	logs := []*model.MgoTrainingLog{}

	skip := getSkipCount(page)
	err = col.Find(query).Skip(skip).Limit(offsetSize).All(&logs)
	if err != nil {
		return nil, err
	}

	return aggregateTrainingLogs(logs)
}

func getSkipCount(page int) int {
	if page <= 0 {
		return 0
	}
	return (page - 1) * offsetSize
}

func aggregateTrainingLogs(rows []*model.MgoTrainingLog) ([]*model.TrainingLog, error) {
	trainings := []*model.TrainingLog{}
	for i := 0; i < len(rows); i++ {
		training := &model.TrainingLog{
			ID:    rows[i].ID.Hex(),
			Date:  rows[i].Date,
			Count: rows[i].Count,
			Kind:  rows[i].Kind,
		}
		trainings = append(trainings, training)
	}
	return trainings, nil
}

func (tp trainingPersistance) InsertTrainingLog(kind string, date string, count int) error {
	sess, err := ConnectMongoDB()
	if err != nil {
		return err
	}
	defer CloseMongoDB(sess)

	col := sess.DB(os.Getenv("MONGO_DBNAME")).C("training_logs")
	item := &model.MgoTrainingLog{
		Kind:  kind,
		Date:  date,
		Count: count,
	}

	return col.Insert(item)
}

func (tp trainingPersistance) UpdateTrainingLog(objectID string, kind string, date string, count int) error {
	sess, err := ConnectMongoDB()
	if err != nil {
		return err
	}
	defer CloseMongoDB(sess)

	col := sess.DB(os.Getenv("MONGO_DBNAME")).C("training_logs")
	query := bson.M{
		"_id": bson.ObjectIdHex(objectID),
	}
	update := &model.MgoTrainingLog{
		Kind:  kind,
		Date:  date,
		Count: count,
	}

	return col.Update(query, update)
}

func (tp trainingPersistance) DeleteTrainingLog(objectID string) error {
	sess, err := ConnectMongoDB()
	if err != nil {
		return err
	}
	defer CloseMongoDB(sess)

	col := sess.DB(os.Getenv("MONGO_DBNAME")).C("training_logs")
	query := bson.M{
		"_id": bson.ObjectIdHex(objectID),
	}

	return col.Remove(query)
}
