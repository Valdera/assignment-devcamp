package product

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	e "shop/internal/entity/product"
)

type ProductRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) AddProduct(ctx context.Context, data e.Product) (result e.Product, err error) {
	var id int64

	if err := r.db.QueryRowContext(ctx, addProductQuery,
		data.Name,
		data.Price,
		data.Description,
		data.Variant,
		data.Discount,
		data.CreatedAt,
		data.UpdatedAt,
	).Scan(&id); err != nil {
		log.Println("[ProductRepo][AddProduct][Storage] problem querying to db, err: ", err.Error())
		return result, err
	}

	result.ID = id

	return
}

func (r *ProductRepo) GetProductAll(ctx context.Context) (result []e.Product, err error) {
	result = make([]e.Product, 0)

	rows, err := r.db.QueryContext(ctx, getProductAllQuery)
	if err != nil {
		log.Println("[ProductRepo][GetProductAll] problem querying to db, err: ", err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var rowData e.Product
		if err = rows.Scan(
			&rowData.ID,
			&rowData.Name,
			&rowData.Price,
			&rowData.Description,
			&rowData.Variant,
			&rowData.Discount,
			&rowData.CreatedAt,
			&rowData.UpdatedAt,
		); err != nil {
			log.Println("[ProductRepo][GetProductAll] problem with scanning db row, err: ", err.Error())
			return
		}
		result = append(result, rowData)
	}

	return
}

func (r *ProductRepo) GetProductById(ctx context.Context, id int64) (result e.Product, err error) {

	rows, err := r.db.QueryContext(ctx, getProductByIdQuery, id)
	if err != nil {
		log.Println("[ProductRepo][GetProductAll] problem querying to db, err: ", err.Error())
		return
	}
	defer rows.Close()

	if rows.Next() {
		if err = rows.Scan(
			&result.ID,
			&result.Name,
			&result.Price,
			&result.Description,
			&result.Variant,
			&result.Discount,
			&result.CreatedAt,
			&result.UpdatedAt,
		); err != nil {
			log.Println("[ProductRepo][GetProductAll] problem with scanning db row, err: ", err.Error())
			return
		}
	} else {
		err = errors.New(fmt.Sprintf("Product with id : %d does not exists", id))
	}

	return
}

func (r *ProductRepo) UpdateProduct(ctx context.Context, id int64, param e.Product) (result e.Product, err error) {
	res, err := r.db.ExecContext(ctx, updateProductQuery,
		param.Name,
		param.Price,
		param.Description,
		param.Variant,
		param.Discount,
		param.CreatedAt,
		param.UpdatedAt,
		id,
	)
	if err != nil {
		log.Println("[ProductRepo][UpdateProduct][Storage] problem querying to db, err: ", err.Error())
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("[ProductRepo][UpdateProduct] problem querying to db, err: ", err.Error())
		return
	}
	if rowsAffected == 0 {
		log.Println("[ProductRepo][UpdateProduct] no rows affected in db")
		return
	}

	result.ID = id

	return
}

func (r *ProductRepo) DeleteProduct(ctx context.Context, id int64) (err error) {
	res, err := r.db.ExecContext(ctx, deleteProductQuery,
		id,
	)
	if err != nil {
		log.Println("[ProductRepo][UpdateProduct][Storage] problem querying to db, err: ", err.Error())
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("[ProductRepo][UpdateProduct] problem querying to db, err: ", err.Error())
		return
	}
	if rowsAffected == 0 {
		log.Println("[ProductRepo][UpdateProduct] no rows affected in db")
		return
	}

	return
}
