package bootstrap

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Application struct {
	DB       *sqlx.DB
	Repos    *Repositories
	Services *Services
	Handlers *Handlers
	Router   http.Handler
}
