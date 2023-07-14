package backend

import (
	"net/http"
)

// indexHandler handles the index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
}
