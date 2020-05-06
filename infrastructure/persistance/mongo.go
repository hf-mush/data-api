package persistance

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"
)

// ConnectMongoDB connect mongodb
func ConnectMongoDB() (*mgo.Session, error) {
	dialInfo := &mgo.DialInfo{
		Addrs: []string{
			os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT"),
		},
		Direct:    true,
		Timeout:   60 * time.Second,
		Database:  os.Getenv("MONGO_DBNAME"),
		Username:  os.Getenv("MONGO_USER"),
		Password:  os.Getenv("MONGO_PASSWORD"),
		Source:    os.Getenv("MONGO_DBNAME"),
		Mechanism: "SCRAM-SHA-1",
	}
	ses, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", "-", err.Error()))
		return ses, err
	}

	limit, err := strconv.Atoi(os.Getenv("MONGO_POOL_LIMIT"))
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", "-", err.Error()))
		return ses, err
	}

	ses.SetPoolLimit(limit)

	if os.Getenv("EXEC_ENV") == "production" {
		ses.SetMode(mgo.PrimaryPreferred, false)
	} else {
		ses.SetMode(mgo.Primary, false)
	}

	return ses, nil
}

// CloseMongoDB close connection mongodb
func CloseMongoDB(session *mgo.Session) {
	if session != nil {
		session.Close()
	}
}
