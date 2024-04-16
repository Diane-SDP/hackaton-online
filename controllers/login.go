package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"hackaton/models"
	"html/template"
	"net/http"
)



func LoginHandler(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("./view/login.html")
	if err != nil {
		panic(err)
	}
	if r.Method == "POST"{


	pseudoL := r.FormValue("pseudo")
	mdpL := r.FormValue("mdp")
	passloghash := sha256.Sum256([]byte(mdpL))
	if pseudoL != "" && mdpL != ""{
		existaccount, psswd, _ := models.ExistAccount(pseudoL)
		if existaccount && psswd == hex.EncodeToString((passloghash[:])) {
			uid := models.Getuid(pseudoL)
			http.SetCookie(w, &http.Cookie{
				Name:   "admin",
				Value:  uid,
				MaxAge: 3600,
			})
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}