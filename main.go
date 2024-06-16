package main

import (
	"fmt"
	"go-postgres-stocksAPI/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Server running on post: 8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}
