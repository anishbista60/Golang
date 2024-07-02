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
        http.Error(w, "method is not supported", http.StatusNotFound)
        return
    }
    fmt.Fprintf(w, "hello!, Nice you make it perfect")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Fprintf(w, "Post request successful done\n")
    name := r.FormValue("name")
    email := r.FormValue("email")
    age := r.FormValue("age")
    gender := r.FormValue("gender")
    message := r.FormValue("message")
    subscribe := r.FormValue("subscribe")

    fmt.Fprintf(w, "Name: %s\n", name)
    fmt.Fprintf(w, "Email: %s\n", email)
    fmt.Fprintf(w, "Age: %s\n", age)
    fmt.Fprintf(w, "Gender: %s\n", gender)
    fmt.Fprintf(w, "Message: %s\n", message)

    subscribeText := "Not Subscribed"
    if subscribe == "on" {
        subscribeText = "Subscribed"
    }
    fmt.Fprintf(w, "Subscribe: %s\n", subscribeText)
}
