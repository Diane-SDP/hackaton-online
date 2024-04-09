package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"html/template"
	"net/http"
	"hackaton/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("./view/login.html")
	if err != nil {
		panic(err)
	}
	if r.Method == "POST"{


	pseudoL := r.FormValue("pseudo")
	mdpL := r.FormValue("mdp")
	pseudoR := r.FormValue("pseudoR")
	mdpR := r.FormValue("mdpR")
	passreghash := sha256.Sum256([]byte(mdpR))
	passloghash := sha256.Sum256([]byte(mdpL))
	if pseudoR != "" && mdpR != ""{
		models.AddUser(pseudoR,hex.EncodeToString(passreghash[:]))
	}
	if pseudoL != "" && mdpL != ""{
		existaccount, psswd, _ := models.ExistAccount(pseudoL)
		if existaccount && psswd == hex.EncodeToString((passloghash[:])) {
			uid := models.Getuid(pseudoL)
			http.SetCookie(w, &http.Cookie{
				Name:   "pseudo_user",
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