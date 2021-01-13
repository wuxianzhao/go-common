package main

import(
	"log"
	"go-common/library/net/http"
)

func main() {
	log.Printf("cms run....")
	http.Serve()
}
