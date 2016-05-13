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
	bind := BindAddr()
	hostname, _ := os.Hostname()

	fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
		fmt.Fprintf(w, "I'm %s\n", hostname)
		for _, whoami := range GetWhoamis() {
			fmt.Fprintf(os.Stdout, "%s: %s\n", whoami.Key, whoami.Value)
			fmt.Fprintf(w, "%s: %s\n", whoami.Key, whoami.Value)
		}
	})

	log.Fatal(http.ListenAndServe(bind, nil))
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

func WhoamiFromEnvStr(txt string) (w *Whoami) {
	if strings.HasPrefix(txt, "WHOAMI") {
		pair := strings.Split(txt, "=")
		key := pair[0][7:len(pair[0])]
		value := pair[1]
		w = &Whoami{key, value}
	}
	return
}
