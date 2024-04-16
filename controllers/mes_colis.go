package controller

import (
	"html/template"
	"net/http"
)

func MesColisHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/mes_colis" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}
	connected := false
	_, err := r.Cookie("admin")
	if err != nil {
		_, err = r.Cookie("shop")
		if err == nil {
			connected = true
		}
	} else {
		connected = true
	}
	tmpl, err := template.ParseFiles("./view/mes_colis.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, connected)
	if err != nil {
		panic(err)
	}
}
