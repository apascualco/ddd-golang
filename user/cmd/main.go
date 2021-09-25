package main

import (
	"log"

	"github.com/apascualco/apascualco-user/cmd/api"
)

func main() {
	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
