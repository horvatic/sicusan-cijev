package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", PipelineServer)
    http.ListenAndServe(":5000", nil)
}

func PipelineServer(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Body)
}
