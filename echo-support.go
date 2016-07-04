package debugcharts

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
)

// EchoDebugRouter registers routes for echo's standard engine
func EchoDebugRouter(e *echo.Echo)  {
	charts := e.Group("/debug/charts")
	{
		charts.GET("/data-feed", standard.WrapHandler(http.HandlerFunc(s.dataFeedHandler)))
		charts.GET("/data", standard.WrapHandler(http.HandlerFunc(dataHandler)))
		charts.GET("/", standard.WrapHandler(http.HandlerFunc(handleAsset("static/index.html"))))
		charts.GET("/main.js", standard.WrapHandler(http.HandlerFunc(handleAsset("static/main.js"))))
		charts.GET("/jquery-2.1.4.min.js", standard.WrapHandler(http.HandlerFunc(handleAsset("static/jquery-2.1.4.min.js"))))
		charts.GET("/moment.min.js", standard.WrapHandler(http.HandlerFunc(handleAsset("static/moment.min.js"))))
	}
}

