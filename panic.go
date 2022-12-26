package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	// The port could be blocked
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}