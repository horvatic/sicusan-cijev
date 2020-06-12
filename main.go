package main

import (
    "net/http"
    "encoding/json"
    "sicusan-cijev/hooks"
    "os"
    "fmt"
    "os/exec"
    "io/ioutil"
)

func main() {
    http.HandleFunc("/", PipelineServer)
    http.ListenAndServe(":5000", nil)
}

func PipelineServer(w http.ResponseWriter, r *http.Request) {
    var g hooks.GitHook

    err := json.NewDecoder(r.Body).Decode(&g)
    if err != nil {
	fmt.Fprintf(w, "Could not read hook")
        return
    }
    if(g.Ref != "refs/heads/master") {
	fmt.Fprintf(w, "Branch was not master branch")
	return
    }


    jsonFile, err := os.Open("approved.json")
    if err != nil {
	fmt.Fprintf(w, "Could not read approved list of repos")
        return
    }
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)
    var approvedRepos []string
    _ = json.Unmarshal([]byte(byteValue), &approvedRepos)

    if(!contains(approvedRepos,g.Repository.URL)) {
	    fmt.Fprintf(w, "Repo not in approved list of repos")
	    return
    }

    cmd := exec.Command("/bin/sh", "release.sh", g.Repository.URL, g.Repository.Name)
    cmd.Run()
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
