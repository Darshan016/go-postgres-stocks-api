package router

import (
	"go-postgres-stocksAPI/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/stocks", middleware.GetStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/stocks/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/stocks/", middleware.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/stocks/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/v1/stocks/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")
	return router
}
