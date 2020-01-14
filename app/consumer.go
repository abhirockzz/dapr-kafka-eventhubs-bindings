package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port string

func init() {
	port = os.Getenv("APP_PORT")
	if port == "" {
		log.Fatalf("missing environment variable %s", "APP_PORT")
	}
}
func main() {
	http.HandleFunc("/timebound", func(rw http.ResponseWriter, req *http.Request) {
		var _time TheTime

		err := json.NewDecoder(req.Body).Decode(&_time)
		if err != nil {
			fmt.Println("error reading message from event hub binding", err)
			rw.WriteHeader(500)
			return
		}
		fmt.Printf("data from Event Hubs '%s'\n", _time)
		rw.WriteHeader(200)
	})
	http.ListenAndServe(":"+port, nil)
}

//TheTime - just for fun(c)
type TheTime struct {
	Time string `json:"time"`
}
