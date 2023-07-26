package internal

import (
	"log"
	"net/http"
	"time"

	"git/ssengerb/ascii-art-web/internal/handler"
)

func Server() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/ascii-art", handler.Ascii)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	server := http.Server{
		Addr:         ":4000",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("The server is running on http://127.0.0.1" + server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
