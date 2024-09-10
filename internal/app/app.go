package app

import (
	"fmt"
	"net/http"

	"bank-system-app/internal/database"
	"bank-system-app/internal/services"
	"bank-system-app/internal/transport"

	"gorm.io/gorm"
)

func NewHTTPServer(mux *http.ServeMux) *http.Server {
	fmt.Println("Starting HTTP server")
	srv := &http.Server{Addr: ":8080", Handler: mux}
	return srv
}

func AddRoutes(db database.Database, mux *http.ServeMux) {
	_ = transport.NewBankHandler(services.NewBankService(db), mux)
}

func Run() {
	var db database.Database
	db, err := database.NewSQLiteDatabase("migrations/main.db", &gorm.Config{})
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	AddRoutes(db, mux)

	server := NewHTTPServer(mux)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
