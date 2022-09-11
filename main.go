package main

import(
    "fmt"
    "net/http"
)

func main (){
    r := new(router)
    r.h = make(map[string]map[string]http.HandlerFunc)

    r.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello!")
    })

    r.HandleFunc("POST", "/", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Bye!")
    })

    http.ListenAndServe(":8080", r)
}
