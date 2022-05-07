package gameStats

import (
	"github.com/zdunker/gameStats/views"
	"github.com/zdunker/webframe"
)

func RunServer() {
	web := webframe.NewEngine()
	dota := web.NewGroup("/dota2")
	dota.GET("/match", views.GetMatch)
	web.Run(":8080")
}
