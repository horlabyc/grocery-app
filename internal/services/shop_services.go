package services

import (
	"context"
	"errors"

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
	existingShop, err := s.repo.GetByID(ct, shop.ID)
	if err != nil {
		return errors.New("shop not found")
	}
	return s.repo.Update(ct, existingShop)
}

// func (s *ShopService) DeleteShop(ct context.Context, id int64) error {
// 	return s.repo.Delete(ct, id)
// }
