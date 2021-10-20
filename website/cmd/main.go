package main

import (
	"log"
	"net/http"

	staff "github.com/MDPaun/go-start/tree/main/website/cmd/account/staff"
)

func main() {

	mux := http.NewServeMux()
	mapUrls()
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func mapUrls() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	staff.StaffUrls()
	http.Handle("/favicon.ico", http.NotFoundHandler())

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello page"))
}
