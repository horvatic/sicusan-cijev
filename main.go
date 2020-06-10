package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "sicusan-cijev/hooks"
)

func main() {
    http.HandleFunc("/", PipelineServer)
    http.ListenAndServe(":5000", nil)
}

func PipelineServer(w http.ResponseWriter, r *http.Request) {
    var g hooks.GitHook

    // Try to decode the request body into the struct. If there is an error,
    // respond to the client with the error message and a 400 status code.
    err := json.NewDecoder(r.Body).Decode(&g)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    fmt.Println(g)
}
