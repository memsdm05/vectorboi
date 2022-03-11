package main

import (
	"vectorboi/app"
	"vectorboi/helpers"
)

func main() {
	helpers.RunGame(new(app.SimpleGame))
}
