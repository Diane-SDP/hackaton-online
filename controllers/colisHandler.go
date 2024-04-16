package controller

import (
	models "hackaton/models"
	"html/template"
	"net/http"
)

type PageColis struct {
	Colis    models.Colis
	Connected bool
}

func ColisHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("colis-id")
	if models.Exist(code) {
		datacolis := models.GetColis(code)
		tmpl, err := template.ParseFiles("./view/mes_colis.html")
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

		// distance := models.GetDistanceFromStart(datacolis.StartCity, datacolis.ClientCity)
		// datacolis.Distance = int(distance)
		// var position int
		// var message string
		// if distance <= 100 {
		// 	position = int(distance / 3)
		// 	message = "Bon toutou , tu fais des efforts ecologiques t'auras un peu de cash back"
		// } else if distance <= 500 {
		// 	message = "Bon ca va , t'es dans la moyenne , t'as ni bonus ni malus"
		// 	position = int(distance*0.08 + 25.21)
		// } else if distance <= 1000 {
		// 	position = int(distance*0.068 + 31)
		// 	message = "T'as commandé tarpin loin frerot fait un effort"
		// } else {
		// 	position = 99
		// 	message = "Fin frerot t'abuses de malade , t'habites a st helene pour commander aussi loin ou quoi ?"
		// }
		// println(position)
		// page := PageColis{
		// 	Colis:    datacolis,
		// 	Position: position,
		// 	Message:  message,
		// }
		page := PageColis {
			Colis: datacolis,
			Connected: connected,
		}
		err = tmpl.Execute(w, page)
		if err != nil {
			panic(err)
		}
	} else {
		NotFound(w, r, http.StatusNotFound)
	}

}
