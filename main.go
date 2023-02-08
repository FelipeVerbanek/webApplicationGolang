package main

import (
	"net/http"

	"github.com/felipeverbanek/app-web/routes"
	_ "github.com/lib/pq"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
