package persistance

import (
	"database/sql"
	"os"
)

var dbPool *sql.DB

// ConnectDB establish mysql connection
func ConnectDB() (*sql.DB, error) {
	cnn, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST")+":"+os.Getenv("MYSQL_PORT")+")/"+os.Getenv("MYSQL_DBNAME"))
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
