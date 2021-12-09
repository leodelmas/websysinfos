package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Démarrage sur le port :8080")
	http.ListenAndServe(":8080", getRouter())
}
