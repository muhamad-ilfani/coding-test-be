package app

import (
	"context"
	"database/sql"
	"log"

	"github.com/labstack/echo"
)

type App struct {
	DB   *sql.DB
	Echo *echo.Echo
}

func Run(ctx context.Context) {
	app := new(App)

	if err := app.initDB(ctx); err != nil {
		log.Fatal(err)
	}

	/*if err := app.initSchema(ctx); err != nil {
		log.Fatal(err)
	}*/

	if err := app.initTable(ctx); err != nil {
		log.Fatal(err)
	}

	app.initService(ctx)
}
