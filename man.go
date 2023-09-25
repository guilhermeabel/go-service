package main

import (
	"github.com/guilhermeabel/go-service/server"
)

func main() {

	app := &server.ApplicationServer{}
	err := app.Server().ListenAndServe()

	if err != nil {
		panic(err)
	}

}
