package controller

import (
	"hackaton/models"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func BOHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./view/BO.html")
	if err != nil {
		panic(err)
	}
	Shop := r.FormValue("shop")
	ClientAdress := r.FormValue("ClientAdress")
	StartAdress := r.FormValue("StartAdress")
	println("les infos : ")
	println(Shop, ClientAdress, StartAdress)
	// var name string
	// cookie, err := r.Cookie("pseudo_user")
	// if err != nil {
	// 	name = ""
	// } else {
	// 	name = cookie.Value
	// }

	// models.GetUserByUid(name)

	if Shop != "" && ClientAdress != "" && StartAdress != "" {
		const charset = "0123456789"
		var seededRand *rand.Rand = rand.New(
			rand.NewSource(time.Now().UnixNano()))
		var code string
		for i := 0; i < 16; i++ {
			code += string(charset[seededRand.Intn(len(charset))])
		}
		idshop, err := strconv.Atoi(Shop)
		if err != nil {
			panic(err)
		}
		models.AddColis(code, idshop, StartAdress, ClientAdress)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
