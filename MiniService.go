package main

import (
        "net/http"
        "log"
        "fmt"
        "net/http/httputil"
)

var sam User
func main() {
        sam = User{"Sam", 12345, "sam@gmai.com", "12345678"}
        server := http.Server{
                Addr:    ":8443",
                Handler: nil,
        }

        http.DefaultServeMux.HandleFunc("/name", nameFunc)
        http.DefaultServeMux.HandleFunc("/id", idFunc)
        http.DefaultServeMux.HandleFunc("/authorize", authzFunc)
        http.DefaultServeMux.HandleFunc("/", allFunc)

        err := server.ListenAndServeTLS("/scratch/openssl/server/newcert.pem", "/scratch/openssl/server/key.pem")
        if err != nil {
           log.Fatal(err)
        }
}
func authzFunc(writer http.ResponseWriter, request *http.Request) {
        log.Printf("Request:\n")
        dump, err := httputil.DumpRequest(request, true)
        if err != nil {
                log.Print(err)
        }
        log.Print(string(dump))

        defer request.Body.Close()

        fmt.Fprint(writer, "{\"apiVersion\": \"authorization.k8s.io/v1beta1\",\"kind\": \"SubjectAccessReview\",\"status\": {\"allowed\":true}}")
    
}


func allFunc(writer http.ResponseWriter, request *http.Request) {
        fmt.Fprint(writer, request.URL.Path)
}

type User struct {
        name string
        id int
        email string
        phone string
}

func idFunc(writer http.ResponseWriter, request *http.Request) {

        fmt.Fprint(writer, sam.id)
}


func nameFunc(w http.ResponseWriter, r *http.Request) {

        fmt.Fprint(w, sam.name)
}
