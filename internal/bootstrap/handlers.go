package bootstrap

import "github.com/horlabyc/grocery-app/internal/handlers"

type Handlers struct {
	Shop *handlers.ShopHandler
}

func InitializeHandlers(services *Services) *Handlers {
	return &Handlers{
		Shop: handlers.NewShopHandler(services.Shop),
	}
}
