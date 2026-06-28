package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
}

var products = []Product{
	{ID: 1, Name: "Laptop Pro", Category: "Electronics", Price: 1299.99, Stock: 45},
	{ID: 2, Name: "Wireless Mouse", Category: "Accessories", Price: 29.99, Stock: 200},
	{ID: 3, Name: "Mechanical Keyboard", Category: "Accessories", Price: 89.99, Stock: 150},
	{ID: 4, Name: "4K Monitor", Category: "Electronics", Price: 599.99, Stock: 30},
	{ID: 5, Name: "USB-C Hub", Category: "Accessories", Price: 49.99, Stock: 300},
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(products)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok","service":"products-go"}`))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/products", productsHandler)
	mux.HandleFunc("/health", healthHandler)

	log.Printf("Go products service listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
