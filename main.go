package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	r.Run(":80")
}
