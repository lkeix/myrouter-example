package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lkeix/myrouter"
)

func main() {
	r := myrouter.NewRouter()
	r.GET("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root"))
	}))

	r.GET("/hoge", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hoge"))
	}))

	r.GET("/:hoge", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		value := myrouter.PathParam(r, "hoge")
		w.Write([]byte("hello: " + value))
	}))

	r.GET("/:hoge/:fuga", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := myrouter.PathParam(r, "hoge")
		f := myrouter.PathParam(r, "fuga")
		fmt.Printf("hello: %s\ndetail: %s\n", h, f)
	}))

	log.Fatal(http.ListenAndServe(":8080", r))
}
