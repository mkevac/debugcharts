// Copyright 2015 mint.zhao.chiu@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package debugcharts

import (
	"github.com/gin-gonic/gin"
	"github.com/mkevac/debugcharts/bindata"
	"log"
)

// add by mint.zhao.chiu@gmail.com
// add supported for gin web framework
func GinDebugRouter(router gin.IRouter) {
	charts := router.Group("/debug/charts")
	{
		charts.GET("/data-feed", ginDataFeedHandler)
		charts.GET("/data", ginDataHandler)
		charts.GET("", ginHandleAsset("static/index.html"))
		charts.GET("/main.js", ginHandleAsset("static/main.js"))
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
