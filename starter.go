package main

//Found in https://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html
import (
	"fmt"
	"net/http"
	"github.com/graphql-go/graphql-go-handler"
	"os"
)

func main() {

	initGraphQl()

	//handler for example price calculation
	http.HandleFunc("/", calculateExamplePrice)

	//handler for GraphQL
	h := handler.New(&handler.Config{
		Schema:   &priceCalculationSchema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("GraphQL-API runs on /graphql")
	fmt.Println("Server is listening ...")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}

}
