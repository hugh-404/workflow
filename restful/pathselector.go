package restful

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/workflow/client"
)

var GetHandlerMap = map[string]gin.HandlerFunc{}
var PostHandlerMap = map[string]gin.HandlerFunc{}

func init() {
	GetHandlerMap["/awx/mock"] = awxMockGetHandler
	PostHandlerMap["/awx/mock"] = awxMockPostHandler
}

func awxMockGetHandler(ctx *gin.Context) {
	params := map[string]interface{}{
		"Header": ctx.Request.Header,
		"Query": ctx.Request.URL.Query(),
	}
	ctx.Request.URL.Query()
	result, err := (&client.AwxCient{}).Run(context.TODO(), params)
	if err != nil {
		ctx.String(500, "error")
	} else {
		ctx.JSON(200, result)
	}
}

func awxMockPostHandler(ctx *gin.Context) {
	(&client.AwxCient{}).Run(context.TODO(), nil)
}