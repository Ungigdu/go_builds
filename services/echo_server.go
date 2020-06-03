package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var echo string  =  "echo"

func Start(port string, s  string)  {
	echo = s
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RespondInJson(echo, w)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func RespondInJson(v interface{}, w http.ResponseWriter) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main(){
	//		port		123
	Start(os.Args[1], os.Args[2])
}

