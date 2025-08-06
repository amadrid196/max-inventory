package repository

import (
	context "context"

	entity "github.com/amadrid196/max-inventory/internal/entity"
)

const (
	qryInsertProduct  = `INSERT INTO products (name, descripcion, price, create_by) VALUES (?, ?, ?, ?)`
	qryGetProducts    = `SELECT id, name, descripcion, price, stock, created_at FROM products`
	qryGetProductByID = `SELECT id, name, descripcion, price, stock, created_at FROM products WHERE id = ?`
)

func (r *repo) SaveProduct(ctx context.Context, name, descripcion string, price float32, createBy int64) error {
	_, err := r.db.ExecContext(ctx, qryInsertProduct, name, descripcion, price, createBy)
	return err
}

func (r *repo) GetProducts(ctx context.Context) ([]entity.Products, error) {
	pp := []entity.Products{}

	err := r.db.SelectContext(ctx, &pp, qryGetProducts)

	if err != nil {
		return nil, err
	}
	return pp, nil
}

func (r *repo) GetProduct(ctx context.Context, id int64) (*entity.Products, error) {
	p := &entity.Products{}

	err := r.db.GetContext(ctx, p, qryGetProductByID, id)

	if err != nil {
		return nil, err
	}
	return p, nil
}
