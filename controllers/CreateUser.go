package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"hackaton/models"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var defaultAdmin bool = true

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./view/Cuser.html")
	if err != nil {
		panic(err)
	}
	if r.Method == "POST" {
		if defaultAdmin {
			models.DeleteDefaultAdmin()
			defaultAdmin = false
		}
		const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
		var mdp string
		for i := 0; i < 10; i++ {
			mdp += string(charset[seededRand.Intn(len(charset))])
		}
		pseudo := r.FormValue("name")
		mail := r.FormValue("mail")
		mdpHash := sha256.Sum256([]byte(mdp))
		models.AddUser(pseudo, hex.EncodeToString(mdpHash[:]),mail)
		models.SendMail(mail,"Vous êtes maintenant administrateur du site :\nPseudo :"+pseudo+"\n\nMot de passe  :"+mdp,"Identifiants Administration TrackTheur")
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
