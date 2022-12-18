package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hello Tour Attraction application!</h1>")
	})

	http.ListenAndServe(":9000", nil)
}
