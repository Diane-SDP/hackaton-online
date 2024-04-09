package controller

import (
    "html/template"
    "net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request, status int) {
    tmpl, err := template.ParseFiles("./view/notFound.html")
    if err != nil {
        panic(err)
    }
    err = tmpl.Execute(w, nil)
    if err != nil {
        panic(err)
    }
}