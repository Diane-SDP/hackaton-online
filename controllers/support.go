package controller

import (
	"html/template"
	"net/http"
)

func SupportHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/support" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}
	tmpl, err := template.ParseFiles("./view/support.html")
	if err != nil {
		panic(err)
	}
	connected := false
	_, err = r.Cookie("admin")
	if err != nil {
		_, err = r.Cookie("shop")
		if err == nil {
			connected = true
		}
	} else {
		connected = true
	}
	err = tmpl.Execute(w, connected)
	if err != nil {
		panic(err)
	}
}
