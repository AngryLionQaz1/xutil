package main

import "net/http"

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hi Girl............."))
	})

	http.HandleFunc("/bye", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("bye bye ,this is v1 httpServer"))
	})

}
