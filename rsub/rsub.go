package main

import (
	"github.com/AndrewGuenther/rpubsub"
)

func main() {
	c := rpubsub.BuildConn()
	defer c.Close()
}
