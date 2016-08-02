package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", lineBotHandler)
	http.ListenAndServe(":9900", nil)
}

func lineBotHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", body)
	w.Write([]byte("OK"))
}
