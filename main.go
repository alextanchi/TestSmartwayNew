package main

import (
	"TestSmartwayNew/cmd/app"
	"log"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}

}
