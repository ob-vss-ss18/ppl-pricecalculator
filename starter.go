package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	port := os.Args[1];
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
	// useless comment
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, world")
}