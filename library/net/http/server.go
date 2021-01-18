package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func ServeGin() (*gin.Engine){

	log.Printf("start server gin....8888")
	r := gin.Default();
	r.Run("0.0.0.0:8888")

	return r
}


//默认http，测试
func ServeHttp()  {

	log.Printf("start server net/http....8888")
	//http.HandleFunc("/ping",func(w http.ResponseWriter,r *http.Request){
	//	io.WriteString(w,"pong")
	//})
	http.ListenAndServe("0.0.0.0:8888",nil)
}

