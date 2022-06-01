package gameStats

import (
	"net/http"

	"github.com/zdunker/gameStats/config"
	"github.com/zdunker/gameStats/views"
	"github.com/zdunker/webframe"
)

func RunServer() {
	web := webframe.NewEngine()
	web.GET("/version", func(c *webframe.Context) {
		c.JSONResponse(http.StatusOK, map[string]interface{}{
			"version": config.GetConfig().GetVersion(),
		})
	})
	dota := web.NewGroup("/dota2")
	dota.GET("/match", views.GetMatch)
	web.Run(":8080")
}
