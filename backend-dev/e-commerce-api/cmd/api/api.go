package api

import (
	"database/sql"
	"e-commerce/internal/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type application struct {
	addr string
	db   *sql.DB
}

func New(addr string, db *sql.DB) *application {
	return &application{
		addr: addr,
		db:   db,
	}
}

func (api *application) Run() error {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("api/v1").Subrouter()

	userHandler := controllers.NewUserHandler()
	userHandler.RegisterRoutes(subRouter)

	log.Println("Server started on port", api.addr)
	return http.ListenAndServe(api.addr, router)
}
