package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/AndrewGuenther/rpubsub"
)

func main() {
	c := rpubsub.BuildConn()
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
