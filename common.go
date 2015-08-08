package rpubsub

import (
	"flag"
	"log"

	"github.com/garyburd/redigo/redis"
)

func BuildConn() (redis.Conn, string) {
	var hostname = flag.String("host", ":6379", "The hostname and port of your Redis instance")
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatal("Please specify a single channel")
	}
	var channel = flag.Args()[0]

	c, err := redis.Dial("tcp", *hostname)
	if err != nil {
		log.Fatal(err)
	}
	return c, channel
}
