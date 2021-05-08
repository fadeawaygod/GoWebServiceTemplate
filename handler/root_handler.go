package handler

import (
	"fmt"
	"net/http"
	"strings"
)

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello")
}
