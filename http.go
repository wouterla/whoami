package main

import (
  "strings"
  "os"
  "fmt"
  "net/http"
  "log"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
 	      fmt.Fprintf(w, "I'm %s\n", hostname)
        fmt.Println()
        for _, e := range os.Environ() {
            pair := strings.Split(e, "=")
            if (strings.HasPrefix("pair[0]", "WHOAMI")) {
              fmt.Println(w, "%s: %s", pair[0], pair[1])
            }
        }

    })

    log.Fatal(http.ListenAndServe(":" + port, nil))
}
