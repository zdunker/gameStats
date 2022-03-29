package readings

import "github.com/zdunker/webframe"

func RunReadingServer() {
	web := webframe.NewEngine()
	readingG := web.NewGroup("/readings")
	readingG.GET("", getReadings)
	readingG.POST("", postReadings)
	web.Run(":8080")
}

func getReadings(ctx *webframe.Context) {

}

func postReadings(ctx *webframe.Context) {

}
