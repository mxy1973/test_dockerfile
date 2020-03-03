package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	fmt.Printf("hello RegisterHandler\n")
	router := httprouter.New()
	router.GET("/", homeHandler)
	router.POST("/", homeHandler)
	return router
}

func main()  {
	hello()
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)
}