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
	var isAdmin bool
	var uid string
	cookie, err := r.Cookie("admin")
	if err != nil {
		cookie, err = r.Cookie("shop")
		if err == nil {
			isAdmin = false
		}
	} else {
		isAdmin = true
		uid = cookie.Value
	}


	if r.Method == "POST" {
		Email := r.FormValue("email")
		ClientAdress := r.FormValue("ClientAdress")
		idShop := r.FormValue("shop")
		StartAdress := r.FormValue("StartAdress")
		ClientCity := r.FormValue("ClientCity")
		StartCity := r.FormValue("StartCity")
		StartPostalCode := r.FormValue("StartPostalCode")
		ClientPostalCode := r.FormValue("ClientPostalCode")
		if ClientAdress != "" && StartAdress != "" && StartCity != "" && ClientCity != "" && StartPostalCode != "" && ClientPostalCode != "" {
			const charset = "0123456789"
			var seededRand *rand.Rand = rand.New(
				rand.NewSource(time.Now().UnixNano()))
			var code string
			for i := 0; i < 10; i++ {
				code += string(charset[seededRand.Intn(len(charset))])
			}
			models.SendMail(Email,"Voici le code de votre colis : "+ code ,"Code colis")
			idint, _ := strconv.Atoi(idShop)
			models.AddColis(code, idint, StartAdress, ClientAdress, StartCity+" "+StartPostalCode, ClientCity+" "+ClientPostalCode,Email)
		} else {
			var data Data
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if data.Type == "plus" {
				idcolis, _ := strconv.Atoi(data.Id)
				_, err = models.DB.Exec("UPDATE colis set step = step + 1 where id = ?", idcolis)
				if err != nil {
					panic(err)
				}
				colis := models.GetColisbyid(idcolis)
				models.SendMail(colis.Mail,"Votre colis avec comme code : "+ colis.Uid+" a avancé" ,"Avancement colis")
				http.Redirect(w, r, "/BO", http.StatusSeeOther)
			} else if data.Type == "moins" {
				idcolis, _ := strconv.Atoi(data.Id)
				_, err = models.DB.Exec("UPDATE colis set step = step - 1 where id = ?", idcolis)
				if err != nil {
					panic(err)
				}
				http.Redirect(w, r, "/BO", http.StatusSeeOther)
			}
		}
		http.Redirect(w, r, "/BO", http.StatusSeeOther)
		return
	}
	tmpl, err := template.ParseFiles("./view/BO.html")
	if err != nil {
		panic(err)
	}
	type Data struct {
		CurrentShop models.Shop
		CurrentAdmin models.User
		IsAdmin bool
		AllColis []models.Colis
		AllShops []models.Shop
	}

	
	var data Data
	data.IsAdmin= isAdmin
	if isAdmin {
		data.CurrentAdmin = models.GetUserByUid(uid)
		data.AllColis =  models.GetAllColis()
		data.AllShops = models.GetAllShops()
	} else {
		data.CurrentShop = models.GetShopByUid(uid)
		data.AllColis = models.GetColisOf(data.CurrentShop.Id)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

