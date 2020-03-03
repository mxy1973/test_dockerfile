package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"html/template"
)

func hello()  {
	fmt.Printf("hello\n")
}



func homeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	t, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Printf("ParseFiles hello.html error: %s", err)
		return
	}
	t.Execute(w, "")
	fmt.Printf("hello homeHandler down\n")
}

