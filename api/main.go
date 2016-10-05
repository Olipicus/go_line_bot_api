package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/line", lineBotHandler)
	http.ListenAndServe(":9900", nil)
}

func lineBotHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	err := ioutil.WriteFile("output.txt", body, 0644)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(body))
}
