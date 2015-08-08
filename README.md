# rpubsub
Simple command-line utilities for Redis [pubsub][pubsub]

## Usage
`rpubsub` creates two executables `rpub` and `rsub` their options are the same
and can be seen below.

```bash
$ r[p|s]ub -h
Usage: r[p|s]ub [--host HOSTNAME:PORT] CHANNEL
  -host=":6379": The hostname and port of your Redis instance
```

`rpub` will publish anything sent to its `STDIN` to the specified `CHANNEL`.

`rsub` will print anything messages it receives on the specified `CHANNEL` to
`STDOUT`.

### Getting the source
`rpubsub` is a properly formatted [Go][golang] project which can be acquired
using the standard

```bash
$ go get github.com/AndrewGuenther/rpubsub
```

If you're unfamiliar with Go, execute the following (make sure you have Go
installed!)

```bash
$ mkdir rpubsub
$ cd rpubsub
$ export GOPATH=$(pwd)
$ go get github.com/AndrewGuenther/rpubsub
$ cd src/github.com/AndrewGuenther/rpubsub
```

### Building the binary
After you've got the source, just `make`!

```bash
$ make
$ cd $GOPATH/bin
$ ls
rpub  rsub
```

[pubsub]: http://redis.io/commands#pubsub
[golang]: https://golang.org/
