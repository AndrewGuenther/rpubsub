package main

import (
	"fmt"
	"os"

	"github.com/AndrewGuenther/rpubsub"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c := rpubsub.BuildConn()
	defer c.Close()

	psc := redis.PubSubConn{c}
	psc.Subscribe(os.Args[1])
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s\n", v.Data)
		case redis.Subscription:
			continue
		case error:
			return
		}
	}
}
