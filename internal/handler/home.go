package handler

import (
	"html/template"
	"log"
	"net/http"

	"git/ssengerb/ascii-art-web/pkg"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		pkg.NotFound(w)
		return
	}

	if r.Method != http.MethodGet {
		pkg.MethodNotAllowed(w)
		return
	}

	tmp, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		pkg.InternalServerError(w)
		return
	}

	err = tmp.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		pkg.InternalServerError(w)
		return
	}
}
