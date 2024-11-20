package main

import (
	"command-line-interface/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("----------Entrypoint----------")

	application := app.Generate()
	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
