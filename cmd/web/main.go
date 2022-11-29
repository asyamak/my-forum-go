package main

import (
	"fmt"
	"log"

	"forum/config"
	app "forum/internal/app"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Init Config Error: %v\n", err)
	}

	fmt.Println("initialisation of NewApp.Start")
	app.NewApp(config).Start() //; err != nil {
	
}
