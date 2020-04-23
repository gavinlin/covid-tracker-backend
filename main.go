package main

import(
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hellow world"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Start server on :5000")
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}