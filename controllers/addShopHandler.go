package controller

import (
	models "hackaton/models"
	"html/template"
	"net/http"
)

func AddShopHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/addShop" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}

	if r.Method == "POST" {
		name := r.FormValue("title")
		mail := r.FormValue("mail")
		adresse := r.FormValue("street")
		ville := r.FormValue("city")
		pays := r.FormValue("state")
		if name != "" && mail != "" && adresse != "" && ville != "" && pays != "" {
			models.AddShop(name, mail, adresse+" "+ville+", "+pays, "")
		}
	}
	tmpl, err := template.ParseFiles("./view/addShop.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
