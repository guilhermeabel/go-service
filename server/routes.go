package server

import (
	"net/http"

	"github.com/alexedwards/flow"
	"github.com/guilhermeabel/go-service/handlers"
)

func (App *ApplicationServer) Routes() http.Handler {

	mux := flow.New()

	mux.Use(App.recoverPanic)

	mux.HandleFunc("/", handlers.Index, http.MethodGet)

	return mux
}
