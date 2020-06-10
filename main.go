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

    err := json.NewDecoder(r.Body).Decode(&g)
    if err != nil {
	fmt.Println("Non push request made")
        return
    }
    fmt.Println(g)
}
