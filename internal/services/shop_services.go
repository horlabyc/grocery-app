package services

import (
	"context"

	"github.com/horlabyc/grocery-app/internal/domain/models"
	"github.com/horlabyc/grocery-app/internal/domain/repositories"
)

type ShopService struct {
	repo repositories.ShopRepository
}

func NewShopService(repo repositories.ShopRepository) *ShopService {
	return &ShopService{
		repo: repo,
	}
}

func (s *ShopService) CreateShop(ct context.Context, shop *models.Shop) error {
	return s.repo.Create(ct, shop)
}

func (s *ShopService) GetShopByID(ct context.Context, id int64) (*models.Shop, error) {
	return s.repo.GetByID(ct, id)
}

func (s *ShopService) GetAllShops(ct context.Context) ([]models.Shop, error) {
	return s.repo.GetAll(ct)
}

func (s *ShopService) UpdateShop(ct context.Context, shop *models.Shop) error {
	return s.repo.Update(ct, shop)
}

func (s *ShopService) DeleteShop(ct context.Context, id int64) error {
	return s.repo.Delete(ct, id)
}
