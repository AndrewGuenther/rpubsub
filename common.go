package rpubsub

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

func BuildConn() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	return c
}
