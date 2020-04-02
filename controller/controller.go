// Package static serves static files like CSS, JavaScript, and images.
package controller

import (
	"net/http"
	"fmt"
	"github.com/puresoul/servgo/router"
)

// Load the routes.
func Load() {
	// Serve static files
	router.Get("/public/*filepath", static)
	router.Get("/", index)
}

// Index maps static files.
func static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
	return
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "index")
	return
}
