package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/AndrewGuenther/rpubsub"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: rpub [--host HOSTNAME:PORT] CHANNEL")
		flag.PrintDefaults()
	}
	c, channel := rpubsub.BuildConn()
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

		c.Do("PUBLISH", channel, line)
	}
}
