package controller

import (
	"html/template"
	"net/http"
)

func RetourHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/retour" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}
	tmpl, err := template.ParseFiles("./view/retour.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
