package main

import "github.com/Caiqm/go-log/log"

func main() {
	log := log.NewLogClient("runtime/log", "log")
	log.Info(`{"a":123}`)
}