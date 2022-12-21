package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/ravshanbek-olimov/Book-store-CRUD/models"
	"github.com/ravshanbek-olimov/Book-store-CRUD/pkg/helper"
)

type OrderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (f *OrderRepo) Create(ctx context.Context, order *models.CreateOrder) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO orders(
			order_id,
			user_id,
			book_id,
			order_price,
			updated_at
		) VALUES ( $1, $2, $3, $4, now() )
	`

	_, err := f.db.Exec(ctx, query,
		id,
		order.UserId,
		order.BookId,
		order.OrderPrice,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *OrderRepo) GetByPKey(ctx context.Context, pkey *models.OrderPrimarKey) (*models.Order, error) {

	var (
		id               sql.NullString
		user_id          sql.NullString
		book_id          sql.NullString
		order_price      sql.NullFloat64
		created_at       sql.NullString
	)

	query := `
		SELECT
			order_id,
			user_id,
			book_id,
			order_price,
			updated_at
		FROM
			orders
		WHERE order_id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&user_id,
			&book_id,
			&order_price,
			&created_at,
		)

	if err != nil {
		return nil, err
	}

	return &models.Order{
		Id:             id.String,
		UserId:        user_id.String,
		BookId:        book_id.String,
		OrderPrice:     order_price.Float64,
		CreatedAt:      created_at.String,
	}, nil
}

func (f *OrderRepo) GetList(ctx context.Context, req *models.GetListOrderRequest) (*models.GetListOrderResponse, error) {

	var (
		resp   = models.GetListOrderResponse{}
		offset = " OFFSET 0"
		limit  = " LIMIT 5"
	)

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			order_id,
			user_id,
			book_id,
			order_price,
			updated_at
		FROM
			orders
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {

		var (      
			id               sql.NullString
			user_id          sql.NullString
			book_id          sql.NullString
			order_price      sql.NullFloat64
			created_at       sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&user_id,
			&book_id,
			&order_price,
			&created_at,
		)

		if err != nil {
			return nil, err
		}

		resp.Orders = append(resp.Orders, &models.Order{
			Id:             id.String,
			UserId:        user_id.String,
			BookId:        book_id.String,
			OrderPrice:     order_price.Float64,
			CreatedAt:      created_at.String,
		})

	}

	return &resp, err
}

func (f *OrderRepo) Update(ctx context.Context, req *models.UpdateOrder) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			orders
		SET
			user_id =     :user_id,
			book_id =     :book_id,
			order_price = :order_price,
			updated_at =  now()
		WHERE order_id =  :order_id
	`

	params = map[string]interface{}{
		"order_id":       req.Id,
		"user_id":        req.Id,
		"book_id":        req.Id,
		"order_price":    req.OrderPrice,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *OrderRepo) Delete(ctx context.Context, req *models.OrderPrimarKey) error {

	_, err := f.db.Exec(ctx, "DELETE FROM orders WHERE order_id = $1", req.Id)
	if err != nil {
		return err
	}

	return err
}