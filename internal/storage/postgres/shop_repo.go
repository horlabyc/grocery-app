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

func (r *ShopRepo) GetByID(ctx context.Context, id int64) (*models.Shop, error) {
	query := `
		SELECT id, name, address, description, contact_phone, created_at, updated_at
		FROM shops
		WHERE id = $1
	`

	shop := &models.Shop{}
	err := r.db.GetContext(ctx, &shop, query, id)
	if err != nil {
		return nil, err
	}
	return shop, nil
}

func (r *ShopRepo) Update(ctx context.Context, shop *models.Shop) error {
	query := `
		UPDATE shops
		SET name = $1, address = $2, description = $3, contact_phone = $4
		WHERE id = $5
		RETURNING updated_at
	`

	row := r.db.QueryRowContext(
		ctx,
		query,
		shop.Name,
		shop.Address,
		shop.Description,
		shop.ContactPhone,
	)
	return row.Scan(&shop.UpdatedAt)
}
