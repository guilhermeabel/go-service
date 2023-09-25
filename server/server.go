package server

import (
	"fmt"
	"net/http"
	"time"
)

type ApplicationServer struct{}

func (app *ApplicationServer) Server() *http.Server {

	return &http.Server{
		Addr:         ":8000",
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}


func (app *ApplicationServer) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				serverError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func serverError(w http.ResponseWriter, err error) {
	message := "The server encountered a problem and could not process your request"
	var data []string
	Error(w, http.StatusInternalServerError, message, data)
}
