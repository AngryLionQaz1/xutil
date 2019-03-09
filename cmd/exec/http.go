package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	httpPost()
}
func httpPost() {

	body_type := "application/json;charset=utf-8"
	resp, err := http.Post("http://127.0.0.1:8086/actuator/shutdown", body_type, nil)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}
