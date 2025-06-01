package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Olá! Seu Web App em Go está rodando no Azure!")
}

func main() {
	http.HandleFunc("/", handler)

	port := ":8080"
	fmt.Println("Servidor rodando na porta", port)
	http.ListenAndServe(port, nil)
}