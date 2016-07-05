package debugcharts

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mkevac/debugcharts/bindata"
)

func GinDebugRouter(router gin.IRouter) {
	charts := router.Group("/debug/charts")
	{
		charts.GET("/data-feed", ginDataFeedHandler)
		charts.GET("/data", ginDataHandler)
		charts.GET("/", ginHandleAsset("static/index.html"))
		charts.GET("/main.js", ginHandleAsset("static/main.js"))
		charts.GET("jquery-2.1.4.min.js", ginHandleAsset("static/jquery-2.1.4.min.js"))
		charts.GET("moment.min.js", ginHandleAsset("static/moment.min.js"))
	}
}

func ginDataFeedHandler(ctx *gin.Context) {
	s.dataFeedHandler(ctx.Writer, ctx.Request)
}

func ginDataHandler(ctx *gin.Context) {
	dataHandler(ctx.Writer, ctx.Request)
}

func ginHandleAsset(path string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := bindata.Asset(path)
		if err != nil {
			log.Fatal(err)
		}

		n, err := ctx.Writer.Write(data)
		if err != nil {
			log.Fatal(err)
		}

		if n != len(data) {
			log.Fatal("wrote less than supposed to")
		}
	}
}
