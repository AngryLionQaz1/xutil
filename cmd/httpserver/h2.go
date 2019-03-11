package main

import (
	"log"
	"net/http"
)

type myHandler struct {
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)
	log.Println("Starting v2 httpserver")
	log.Fatal(http.ListenAndServe(":8086", mux))
}

func sayBye(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("bye bye ,this is v2 httpServer"))
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is version 2"))
}
