package main

import (
	"apitest/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("welcome to my api")
	fmt.Println("server is getting starting...")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
}


