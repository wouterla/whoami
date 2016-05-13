package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Whoami struct {
	Key   string
	Value string
}

func main() {
	port := GetPort()
	fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
	hostname, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
		fmt.Fprintf(w, "I'm %s\n", hostname)
		fmt.Println()
		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			if strings.HasPrefix("pair[0]", "WHOAMI") {
				fmt.Fprintf(os.Stdout, "%s: %s", pair[0], pair[1])
				fmt.Fprintf(w, "%s: %s", pair[0], pair[1])
				fmt.Println()
			}
		}

	})

	log.Fatal(http.ListenAndServe(BindAddr(), nil))
}

func GetPort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return
}

func BindAddr() string {
	return ":" + GetPort()
}

func GetWhoamis() (whoamis []*Whoami) {
	for _, kv := range os.Environ() {
		if whoami := WhoamiFromEnvStr(kv); whoami != nil {
			whoamis = append(whoamis, whoami)
		}
	}
	return
}

func WhoamiFromEnvStr(txt string) *Whoami {
	if strings.HasPrefix(txt, "WHOAMI") {
		pair := strings.Split(txt, "=")
		key := pair[0][7:len(pair[0])]
		value := pair[1]
		return &Whoami{key, value}
	}
	return nil
}
