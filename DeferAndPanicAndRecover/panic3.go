package main

import (
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go"))
	})
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err.Error())
	}
}
