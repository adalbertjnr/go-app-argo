package main

import (
	"go-app/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":5000", nil)
}
