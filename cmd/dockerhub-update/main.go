package main

import (
	"log"
	"os"

	update "github.com/charliekenney23/dockerhub-update"
)

func main() {
	if err := update.Entrypoint.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
