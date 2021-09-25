package main

import (
	"log"
	"os"
	"strconv"

	"github.com/apascualco/apascualco-auth/cmd/api"
	"github.com/apascualco/apascualco-auth/cmd/worker"
)

func main() {
	isWorkerEnv := os.Getenv("WORKER")
	isWorker, err := strconv.ParseBool(isWorkerEnv)
	if err == nil && isWorker {
		if err := worker.Start(); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := api.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
