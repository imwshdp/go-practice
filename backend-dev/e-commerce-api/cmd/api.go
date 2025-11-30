package main

import (
	repo "e-commerce-api/internal/adapters/postgresql/sqlc"
	"e-commerce-api/internal/orders"
	"e-commerce-api/internal/products"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

type application struct {
	config config
	db     *pgx.Conn
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	// di
	repository := repo.New(app.db)

	productService := products.NewService(repository)
	productHandler := products.NewHandler(productService)

	r.Get("/products", productHandler.ListProducts)

	orderService := orders.NewService(repository, app.db)
	orderHandler := orders.NewHandler(orderService)

	r.Post("/orders", orderHandler.PlaceOrder)

	return r
}

func (app *application) run(h http.Handler) error {
	srv := http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf(`Server started on port %s`, app.config.addr)

	return srv.ListenAndServe()
}
