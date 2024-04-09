package controller

import (
	"html/template"
	"net/http"
)

func SupportHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./view/support.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
