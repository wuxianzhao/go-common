package http

import "github.com/gin-gonic/gin"

func Serve() (e *gin.Engine,err error) {

	r := gin.Default();
	r.Run("0.0.0.0:8888")

	return
}