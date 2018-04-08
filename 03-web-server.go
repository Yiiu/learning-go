package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/test", testRequest)
	http.HandleFunc("/", getRequest)
	fmt.Println("Linstening on 9999")
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func testRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, r)
}
func getRequest(w http.ResponseWriter, r *http.Request) {
	fileRequested := "./" + r.URL.Path
	http.ServeFile(w, r, fileRequested)
	return
}