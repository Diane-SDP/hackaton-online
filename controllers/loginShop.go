package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"hackaton/models"
	"html/template"
	"net/http"
)

func LoginShopHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./view/loginShop.html")
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
	if r.Method == "POST" {
		mail := r.FormValue("mail")
		mdp := r.FormValue("mdp")
		passloghash := sha256.Sum256([]byte(mdp))
		if mail != "" && mdp != "" {
			existaccount, psswd, uid := models.ExistShop(mail)
			if existaccount && psswd == hex.EncodeToString((passloghash[:])) {
				http.SetCookie(w, &http.Cookie{
					Name:   "shop",
					Value:  uid,
					MaxAge: 3600,
				})
				http.SetCookie(w, &http.Cookie{
					Name:   "admin",
					MaxAge: -1,
				})
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		}
	}
	err = tmpl.Execute(w, connected)
	if err != nil {
		panic(err)
	}
}
