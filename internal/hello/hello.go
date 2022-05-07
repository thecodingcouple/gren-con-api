package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type hello struct {
	Greeting string `json:"greeting"`
	ApiName  string `json:"apiName"`
}

func HelloWorld(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, hello{"Hello World, this API is for", "Gren Con"})
}
