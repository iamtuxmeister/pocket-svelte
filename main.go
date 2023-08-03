package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {

	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Static("/", "build")

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

}
