package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	fmt.Println("DSN: " + config.DatabaseDSN)
	r := router.Generate()

	address := fmt.Sprintf(":%d", config.ApiPort)
	fmt.Println(address)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), r))
}
