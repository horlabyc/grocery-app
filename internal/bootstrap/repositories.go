package bootstrap

import (
	"github.com/horlabyc/grocery-app/internal/domain/repositories"
	"github.com/horlabyc/grocery-app/internal/storage/postgres"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Shop repositories.ShopRepository
}

func InitializeRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Shop: postgres.NewShopRepo(db),
	}
}
