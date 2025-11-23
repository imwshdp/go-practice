package api

import (
	"database/sql"
	"log"
	"net/http"
	cHttp "rest/internal/handlers/cart"
	pHttp "rest/internal/handlers/product"
	uHttp "rest/internal/handlers/user"

	oStorage "rest/internal/storage/postgres/order"
	pStorage "rest/internal/storage/postgres/product"
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

	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// repos
	userRepo := uStorage.NewUserRepository(app.db)
	productsRepo := pStorage.NewProductRepository(app.db)
	orderRepo := oStorage.NewOrderRepository(app.db)

	// handlers
	userHandler := uHttp.NewUserHandler(userRepo)
	userHandler.RegisterRoutes(subRouter)

	productHandler := pHttp.NewProductHandler(productsRepo)
	productHandler.RegisterRoutes(subRouter)

	cartHandler := cHttp.NewOrderHandler(orderRepo, productsRepo, userRepo)
	cartHandler.RegisterRoutes(subRouter)

	log.Println("Server started on port", app.addr)
	return http.ListenAndServe(app.addr, router)
}
