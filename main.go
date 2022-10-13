package main

import (
	"log"
	"net/http"
	"shop/database"
	"shop/server"
	"time"

	"github.com/gorilla/mux"

	productHandler "shop/internal/handler/http/product"
	productRepo "shop/internal/repo/product"
	productUc "shop/internal/usecase/product"
)

func main() {
	dbConfig := database.Config{
		User:     "postgres",
		Password: "12345",
		DBName:   "devcamp",
		Port:     5432,
		Host:     "db",
		SSLMode:  "disable",
	}

	db := database.GetDatabaseConnection(dbConfig)

	productRepo := productRepo.New(db)
	productUc := productUc.New(productRepo)
	productHandler := productHandler.New(productUc)

	router := mux.NewRouter()

	router.HandleFunc("/product/{id}", productHandler.UpdateProductHandler).Methods(http.MethodPatch)
	router.HandleFunc("/product/{id}", productHandler.GetProductHandler).Methods(http.MethodGet)
	router.HandleFunc("/product/{id}", productHandler.DeleteProductHandler).Methods(http.MethodDelete)
	router.HandleFunc("/products", productHandler.GetProductAllHandler).Methods(http.MethodGet)
	router.HandleFunc("/", productHandler.RootHandler).Methods(http.MethodGet)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         9090,
	}
	log.Println("Devcamp-2022 product service service is starting...")

	server.Serve(serverConfig, router)

}
