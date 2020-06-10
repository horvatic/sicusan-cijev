package main

import (
    "net/http"
    "encoding/json"
    "sicusan-cijev/hooks"
    "os/exec"
)

func main() {
    http.HandleFunc("/", PipelineServer)
    http.ListenAndServe(":5000", nil)
}

func PipelineServer(w http.ResponseWriter, r *http.Request) {
    var g hooks.GitHook

    err := json.NewDecoder(r.Body).Decode(&g)
    if err != nil {
        return
    }
    cmd := exec.Command("/bin/sh", "release.sh", g.Repository.URL, g.Repository.Name)
    cmd.Run()
}
