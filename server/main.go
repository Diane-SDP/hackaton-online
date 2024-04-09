package main

import (
	"fmt"
	controller "hackaton/controllers"
	"hackaton/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	models.OpenDb()
	err := models.DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the MySQL database")
	models.CreateDB()
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", controller.HomeHandler)
	http.HandleFunc("/colis", controller.ColisHandler)
	http.HandleFunc("/support", controller.SupportHandler)
	http.HandleFunc("/BO", controller.BOHandler)
	http.HandleFunc("/info", controller.InfoHandler)
	http.HandleFunc("/mes_colis", controller.MesColisHandler)
	http.HandleFunc("/retour", controller.RetourHandler)
	http.HandleFunc("/login", controller.LoginHandler)

	panic(http.ListenAndServe(":8080", nil))

}
