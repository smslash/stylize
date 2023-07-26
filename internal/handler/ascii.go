package handler

import (
	"html/template"
	"log"
	"net/http"

	"git/ssengerb/ascii-art-web/pkg"
)

func Ascii(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.MethodNotAllowed(w)
		return
	}

	tmp, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		log.Println(err.Error())
		pkg.InternalServerError(w)
		return
	}

	input := r.FormValue("convert")
	if len(input) > 300 {
		pkg.BadRequest(w)
		return
	}

	banner := r.FormValue("fonts")
	result, number, err := pkg.Font(banner, input)
	if err != nil {
		if number == 2 {
			pkg.InternalServerError(w)
		} else {
			pkg.BadRequest(w)
		}
		return
	}

	err = tmp.Execute(w, result)
	if err != nil {
		log.Println(err.Error())
		pkg.InternalServerError(w)
		return
	}
}
