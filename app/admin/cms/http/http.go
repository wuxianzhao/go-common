package http

import (
	"go-common/library/net/http"
	"github.com/gin-gonic/gin"
)

func Init() {
	e := http.ServeGin()
	e.router()
}

func router()  {
	
}