package main

import (
	"log"
	"net/http"

	rice "github.com/giter/go.rice"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)
func main() {

	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// e.Router.Static("/", "build")
		assetHandler := http.FileServer(rice.MustFindBox("build").HTTPBox())
		e.Router.GET("/*", echo.WrapHandler(assetHandler))

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

}
