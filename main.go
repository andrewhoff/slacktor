package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var token = os.Getenv("TOKEN")

type hackerQuote struct {
	Quote string `json:"quote"`
}

func handleCommand(w http.ResponseWriter, r *http.Request) {
	if token != r.FormValue("token") {
		log.Println("Token mismatch")
		http.Error(w, "Token mismatch", http.StatusBadRequest)
		return
	}

	resp, err := http.Get("https://hacker.actor/quote")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	q := &hackerQuote{}
	jdec := json.NewDecoder(resp.Body)
	err = jdec.Decode(q)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(q.Quote))
}

func main() {
	http.HandleFunc("/", handleCommand)      // set router
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
