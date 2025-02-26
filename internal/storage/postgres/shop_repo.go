package postgres

import (
	"context"
	"database/sql"
	"errors"

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
		SELECT id, name, description, contact_phone, address, created_at, updated_at
		FROM shops
		WHERE id = $1
	`

	shop := &models.Shop{}
	err := r.db.GetContext(ctx, shop, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("shop not found")
		}
		return nil, err
	}

	return shop, nil
}

func (r *ShopRepo) Update(ctx context.Context, shop *models.Shop) error {
	query := `
		UPDATE shops
		SET name = $1, address = $2, description = $3, contact_phone = $4, updated_at = NOW()
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
		shop.ID,
	)
	return row.Scan(&shop.UpdatedAt)
}

func (r *ShopRepo) Delete(ctx context.Context, id int64) error {
	// First check if this shop is referenced by any grocery items
	var count int
	err := r.db.GetContext(
		ctx,
		&count,
		"SELECT COUNT(*) FROM grocery_items WHERE shop_id = $1",
		id,
	)

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("cannot delete shop that is referenced by grocery items")
	}

	_, err = r.db.ExecContext(ctx, "DELETE FROM shops WHERE id = $1", id)
	return err
}
