package main
//Found in https://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html
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
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, world")
}