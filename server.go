package main

import (
	"github.com/gin-gonic/gin"
	"github.com/workflow/restful"
)

func main() {
	r := gin.Default()
	for path, f := range restful.GetHandlerMap {
		r.GET(path, f)
	}
	for path, f := range restful.PostHandlerMap {
		r.POST(path, f)
	}
	r.Run(":8081")
}
