package persistance

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

var cachePool *redis.Pool

// GetRedisPool return redis connection pool
func GetRedisPool() (*redis.Pool, error) {
	// コネクションの最大数を取得する
	maxConnSize, err := strconv.Atoi(os.Getenv("REDIS_MAX_CONN"))
	if err != nil {
		log.Println("error: [redis] convert int " + err.Error())
	}
	// コネクションプールの最大数を取得する
	maxPoolSize, _ := strconv.Atoi(os.Getenv("REDIS_MAX_CONN_POOL"))
	if err != nil {
		log.Println("error: [redis] convert int " + err.Error())
	}
	// コネクションプールを生成する
	cachePool = &redis.Pool{
		Wait:        true,
		IdleTimeout: 30 * time.Second,
		MaxActive:   maxConnSize,
		MaxIdle:     maxPoolSize,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(
				"tcp",
				os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
	return cachePool, nil
}

// GetConn return redis pool connection
func GetConn() redis.Conn {
	return cachePool.Get()
}

// RedisGet get key/value
func RedisGet(key string) (string, error) {
	// Redis のコネクションを取得
	conn := GetConn()
	defer conn.Close()
	// Redis から保持している key のデータを取得する
	data, err := redis.String(conn.Do("GET", key))
	if err != nil && err != redis.ErrNil {
		return "", err
	}
	return data, nil
}

// RedisSetJSON convert to json
func RedisSetJSON(key string, data interface{}) error {
	// mapをjsonに変換する
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// Redis のコネクションを取得
	conn := GetConn()
	defer conn.Close()
	// json を Redis に保存する
	_, err = conn.Do("SET", key, string(b))
	return err
}

// RedisSet save key/value
func RedisSet(key string, data string) error {
	conn := GetConn()
	defer conn.Close()
	_, err := conn.Do("SET", key, data)
	return err
}
