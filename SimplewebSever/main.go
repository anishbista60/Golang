package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    fileServer := http.FileServer(http.Dir("./Template"))
    http.Handle("/", fileServer)
    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/hello", helloHandler)

    fmt.Printf("Starting server at port 8081\n")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatal(err)
    }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found", http.StatusNotFound)
        return
    }
    if r.Method != "GET" {
        http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
        return
    }
    fmt.Fprintf(w, "hello! Everyone Hope you all are doing well")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Fprintf(w, "Post request successful\n")
    
    firstname := r.FormValue("firstname")
    lastname := r.FormValue("lastname")
    
    fmt.Fprintf(w, "First Name: %s\n", firstname)
    fmt.Fprintf(w, "Last Name: %s\n", lastname)
}

