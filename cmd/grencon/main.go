package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thecodingcouple/gren-con-api/internal/hello"
)

func main() {
	router := gin.Default()
	router.GET("/hello", hello.HelloWorld)
	router.Run("localhost:8080")
}
