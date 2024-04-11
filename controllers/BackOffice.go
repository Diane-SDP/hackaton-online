package controller

import (
	"encoding/json"
	"hackaton/models"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Data struct {
	Type string `json:"type"`
	Mess string `json:"mess"`
	Id   string `json:"id"`
}

func BOHandler(w http.ResponseWriter, r *http.Request) {
	var name string
	cookie, err := r.Cookie("pseudo_user")
	if err != nil {
		name = ""
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		name = cookie.Value
	}

	user := models.GetUserByUid(name)
	if user.Admin != 1 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		ClientAdress := r.FormValue("ClientAdress")
		StartAdress := r.FormValue("StartAdress")
		ClientCity := r.FormValue("ClientCity")
		StartCity := r.FormValue("StartCity")
		StartPostalCode := r.FormValue("StartPostalCode")
		ClientPostalCode := r.FormValue("ClientPostalCode")
		plus := r.FormValue("p")
		moins := r.FormValue("m")
		println("--------")
		println(plus)
		println(moins)
		println("--------")
		if ClientAdress != "" && StartAdress != "" && StartCity != "" && ClientCity != "" && StartPostalCode != "" && ClientPostalCode != "" {
			const charset = "0123456789"
			var seededRand *rand.Rand = rand.New(
				rand.NewSource(time.Now().UnixNano()))
			var code string
			for i := 0; i < 10; i++ {
				code += string(charset[seededRand.Intn(len(charset))])
			}
			models.AddColis(code, 0, StartAdress, ClientAdress, StartCity+" "+StartPostalCode, ClientCity+" "+ClientPostalCode)
		} else {
			var data Data
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			println("giga pute")
			if data.Type == "plus" {
				idcolis, _ := strconv.Atoi(data.Id)
				println("plus")
				_, err = models.DB.Exec("UPDATE colis set step = step + 1 where id = ?", idcolis)
				if err != nil {
					panic(err)
				}
				http.Redirect(w, r, "/BO", http.StatusSeeOther)
			} else if data.Type == "moins" {
				println("moins")
				idcolis, _ := strconv.Atoi(data.Id)
				_, err = models.DB.Exec("UPDATE colis set step = step - 1 where id = ?", idcolis)
				if err != nil {
					panic(err)
				}
				http.Redirect(w, r, "/BO", http.StatusSeeOther)
			}
		}

	}
	tmpl, err := template.ParseFiles("./view/BO.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, models.GetAllColis())
	if err != nil {
		panic(err)
	}
}
