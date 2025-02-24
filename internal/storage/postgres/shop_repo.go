package postgres

import (
	"context"

	"github.com/horlabyc/grocery-app/internal/domain/models"
	"github.com/jmoiron/sqlx"
)

type ShopRepo struct {
	db *sqlx.DB
}

func NewShopRepo(db *sqlx.DB) *ShopRepo {
	return &ShopRepo{
		db: db,
	}
}

func (r *ShopRepo) Create(ctx context.Context, shop *models.Shop) error {
	query := `
		INSERT INTO shops (name, address, description, contact_phone)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	row := r.db.QueryRowContext(
		ctx,
		query,
		shop.Name,
		shop.Address,
		shop.Description,
		shop.ContactPhone,
	)

	return row.Scan(&shop.ID, &shop.CreatedAt, &shop.UpdatedAt)
}

func (r *ShopRepo) GetAll(ctx context.Context) ([]models.Shop, error) {
	query := `
		SELECT id, name, address, description, contact_phone, created_at, updated_at
		FROM shops
		ORDER BY created_at DESC
	`
	var shops []models.Shop
	err := r.db.SelectContext(ctx, &shops, query)
	if err != nil {
		return nil, err
	}
	return shops, nil
}
