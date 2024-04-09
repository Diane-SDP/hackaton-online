package controller

import (
	models "hackaton/models"
	"html/template"
	"net/http"
)

func ColisHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("colis-id")
	if models.Exist(code) {
		datacolis := models.GetColis(code)
		tmpl, err := template.ParseFiles("./view/mes_colis.html")
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(w, datacolis)
		if err != nil {
			panic(err)
		}
	} else {
		NotFound(w, r, http.StatusNotFound)
	}

}
