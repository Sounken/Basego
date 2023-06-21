package main

import (
	"fmt"
	"net/http"

	db "github.com/Sounken/Basego/internal/database"
	forum "github.com/Sounken/Basego/internal/handlers"
)

const port = ":8080"

func main() {
	db.Create()

	// Définir le répertoire où se trouvent les fichiers statiques
	fs := http.FileServer(http.Dir("../../template"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", forum.Home)

	fmt.Print("(http://localhost:8080) - Server running on port ", port ,"\n")
	http.ListenAndServe(port, nil)
}
