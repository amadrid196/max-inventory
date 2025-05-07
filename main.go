package main

import (
	"context"

	"github.com/amadrid196/max-inventory/database"
	"github.com/amadrid196/max-inventory/internal/repository"
	"github.com/amadrid196/max-inventory/internal/service"
	"github.com/amadrid196/max-inventory/settings"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
		),
		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("SELECT * FROM users")
				if err != nil {
					panic(err)
				}
			},
		),
	)
	app.Run()
}
