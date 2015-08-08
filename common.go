package rpubsub

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
)

func BuildConn() (redis.Conn, string) {
	var hostname = flag.String("host", ":6379", "The hostname and port of your Redis instance")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Please specify a single channel.")
		os.Exit(1)
	}
	var channel = flag.Args()[0]

	c, err := redis.Dial("tcp", *hostname)
	if err != nil {
		log.Fatal(err)
	}
	return c, channel
}
