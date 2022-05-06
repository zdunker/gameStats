package main

import (
	"github.com/zdunker/gameStats"
	"github.com/zdunker/gameStats/config"
)

func main() {
	config.LoadConfig(".")
	gameStats.RunServer()
}
