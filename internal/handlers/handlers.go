package forum

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "../../internal/database/db.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tmpl, err := template.ParseFiles("../../template/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//create cookie

	if r.Method == http.MethodGet {
		tmpl.Execute(w, r)
	}
}
