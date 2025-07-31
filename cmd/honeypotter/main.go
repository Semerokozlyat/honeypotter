package main

import "github.com/Semerokozlyat/honeypotter/internal"

func main() {
	app := internal.NewApp()

	if err := app.Start(); err != nil {
		panic(err)
	}
}
