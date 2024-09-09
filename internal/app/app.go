package app

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"bank-system-app/internal/database"
	"bank-system-app/internal/services"
	"bank-system-app/internal/transport"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
	fmt.Println("Starting HTTP server")
	srv := &http.Server{Addr: ":8080", Handler: mux}
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				ln, err := net.Listen("tcp", srv.Addr)
				if err != nil {
					return err
				}
				fmt.Println("Starting HTTP server at", srv.Addr)
				go func() {
					err := srv.Serve(ln)
					if err != nil {
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return srv.Shutdown(ctx)
			},
		},
	)
	return srv
}

func AddRoutes(db database.Database, mux *http.ServeMux) fx.Option {
	return fx.Module(
		"add_routes",
	)
}

func Run() {
	fx.New(
		fx.Provide(
			func() *gorm.Config {
				return &gorm.Config{}
			},
		),
		fx.Provide(
			func() string { return "migrations/main.db" },
			fx.Annotate(
				database.NewSQLiteDatabase,
				fx.As(new(database.Database)),
			),
		),
		fx.Provide(func(database database.Database) *gorm.DB { return database.GetConnection() }),
		fx.Provide(http.NewServeMux),
		// fx.Provide(NewHTTPServer),
		// fx.Invoke(func(server *http.Server) {}),
		fx.Provide(services.NewBankService),
		fx.Provide(transport.NewBankHandler),
		fx.Invoke(func(handler transport.BankHandler) {}),
		fx.Provide(NewHTTPServer),
		fx.Invoke(func(server *http.Server) {}),
	).Run()
}
