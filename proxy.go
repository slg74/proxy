package main

import (
        "fmt"
        "log"
        "net/http"
        )

func main() {

        tr := &http.Transport{
                DisableCompression: true,
                DisableKeepAlives: false,
        }
        client := &http.Client{Transport: tr}

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

                send(r, client)
                fmt.Fprintf(w, "OK")
        })
        log.Fatal(http.ListenAndServe(":8080", nil))
}

func send(r *http.Request, client *http.Client) int {
        req, err := http.NewRequest("GET", "http://localhost:80" + r.URL.Path, nil)
        if err != nil {
                log.Fatal(err)
                return 0
        }       
        resp, err := client.Do(req)
        if err != nil {
                log.Fatal(err)
                return 0
        }       
        if resp == nil {
                return 0
        }       
        return 1
}       
