package backend

import (
	"net/http"
)

// StartFileServers starts the file servers
func StartFileServers() {
	//css fileserver
	fs := http.FileServer(http.Dir("html/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	//js fileserver
	fs = http.FileServer(http.Dir("html/js"))
	http.Handle("/js/", http.StripPrefix("/js/", fs))
}

// StartHandlers starts the handlers
func StartHandlers() {
	http.HandleFunc("/", indexHandler)
}

// StartServer starts the server on port 8080
func StartServer() {
	http.ListenAndServe(":8080", nil)
}
