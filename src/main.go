package main

import (
	"goServer/src/logger"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	logger.Debug_MSG("server close down(signal: %v)", sig)
}
