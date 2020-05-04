package persistance

import (
	"database/sql"
	"os"
	"time"
)

var dbPool *sql.DB

// ConnectDB establish mysql connection
func ConnectDB() (*sql.DB, error) {
	cnn, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST")+":"+os.Getenv("MYSQL_PORT")+")/"+os.Getenv("MYSQL_DBNAME")+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	cnn.SetMaxOpenConns(50)
	cnn.SetMaxIdleConns(10)

	dbPool = cnn

	return cnn, nil
}

// GetConn get mysql connection
func GetConn() *sql.DB {
	return dbPool
}

func convertTimeToJstStr(date time.Time) string {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	return date.In(jst).Format(time.RFC3339)
}
