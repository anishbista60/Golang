package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anishbista60/Golang/book-store/pkg/routes"
	"github.com/gorilla/mux"
)

func main()  {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	fmt.Println("Starting the server at port 9010:")
	log.Fatal(http.ListenAndServe("localhost:9010",r))
}