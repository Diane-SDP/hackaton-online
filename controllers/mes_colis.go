package controller

import (
	"html/template"
	"net/http"
)

func MesColisHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./view/mes_colis.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
