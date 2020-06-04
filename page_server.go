package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func Start()  {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	RespondInJson(echo, w)
	//})

	log.Fatal(http.ListenAndServe(":80", nil))
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
	Start()
}

