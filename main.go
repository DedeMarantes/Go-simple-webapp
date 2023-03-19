package main

import (
	"modulo/routes"
	"net/http"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":8000", nil)
}
