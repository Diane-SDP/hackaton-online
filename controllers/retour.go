package controller

import (
	"html/template"
	"net/http"
)

func RetourHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./view/retour.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
