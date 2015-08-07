package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}

		c.Do("PUBLISH", os.Args[1], line)
	}
}
