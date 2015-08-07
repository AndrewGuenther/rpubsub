package main

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
}
