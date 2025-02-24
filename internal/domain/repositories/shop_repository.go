package repositories

import (
	"context"

	"github.com/horlabyc/grocery-app/internal/domain/models"
)

type ShopRepository interface {
	Create(ctx context.Context, shop *models.Shop) error
	GetAll(ctx context.Context) ([]models.Shop, error)
	// GetByID(cxt context.Context, id int64) (*models.Shop, error)
	// Update(ctx context.Context, shop *models.Shop) error
	// Delete(ctx context.Context, id int64) error
}
