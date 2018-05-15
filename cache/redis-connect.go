package cache

import (
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
)

var db *redis.Pool

func newPool(hostAndPort string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 1000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", hostAndPort)
			if err != nil {
				log.Printf("Connect redis failure:%s\n", err)
				os.Exit(1)
			}
			return c, err
		},
	}
}

// InitConnect initialize cache pool for using later
func InitConnect(hostAndPort string) (*redis.Pool, error) {
	db = newPool(hostAndPort)
	c := db.Get()
	defer c.Close()
	err := c.Err()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetDBHandler get one connection from pool
func GetDBHandler() redis.Conn {
	return db.Get()
}
