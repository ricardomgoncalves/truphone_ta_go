package main

import (
	"github.com/ricardomgoncalves/truphone_ta_go/cmd/service/app"
	"log"
)

func main() {
	if err := app.New(); err != nil {
		log.Fatalln(err.Error())
	}
}
