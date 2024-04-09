package controller

import (
	"html/template"
	"net/http"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./view/info.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
