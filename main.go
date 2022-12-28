package main

import (
	"attrtour/routes"
	"net/http"
)

func init() {
	api := routes.Api{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.Run(w, r)
	})
}

func main() {
	http.ListenAndServe(":9000", nil)
}
