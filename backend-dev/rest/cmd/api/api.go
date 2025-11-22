package api

import (
	"database/sql"
	"log"
	"net/http"
	uHttp "rest/internal/handlers/user"
	uStorage "rest/internal/storage/postgres/user"

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

func (app *application) Run() error {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("api/v1").Subrouter()

	// repos
	userRepo := uStorage.NewUserRepository(app.db)

	// handlers
	userHandler := uHttp.NewUserHandler(userRepo)
	userHandler.RegisterRoutes(subRouter)

	log.Println("Server started on port", app.addr)
	return http.ListenAndServe(app.addr, router)
}
