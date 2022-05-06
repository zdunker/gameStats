package gameStats

import (
	"github.com/zdunker/gameStats/views"
	"github.com/zdunker/webframe"
)

func RunServer() {
	web := webframe.NewEngine()
	// readingG := web.NewGroup("/readings")
	// readingG.GET("", views.GetReadings)
	// readingG.POST("", views.PostReadings)
	dota := web.NewGroup("/dota2")
	dota.GET("/match", views.GetMatch)
	web.Run(":8080")
}
