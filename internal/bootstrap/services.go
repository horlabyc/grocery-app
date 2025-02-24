package bootstrap

import "github.com/horlabyc/grocery-app/internal/services"

type Services struct {
	Shop *services.ShopService
}

func InitializeServices(repos *Repositories) *Services {
	return &Services{
		Shop: services.NewShopService(repos.Shop),
	}
}
