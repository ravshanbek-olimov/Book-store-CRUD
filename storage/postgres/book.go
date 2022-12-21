
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

type BookRepo struct {
	db *pgxpool.Pool
}

func NewBookRepo(db *pgxpool.Pool) *BookRepo {
	return &BookRepo{
		db: db,
	}
}

func (f *BookRepo) Create(ctx context.Context, book *models.CreateBook) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO books(
			book_id,
			book_name,
			author_name,
			book_price,
			updated_at
		) VALUES ( $1, $2, $3, $4, now() )
	`

	_, err := f.db.Exec(ctx, query,
		id,
		book.BookName,
		book.AuthorName,
		book.BookPrice,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *BookRepo) GetByPKey(ctx context.Context, pkey *models.BookPrimarKey) (*models.Book, error) {

	var (
		id            sql.NullString
		book_name     sql.NullString
		author_name   sql.NullString
		book_price    sql.NullFloat64
		createdAt     sql.NullString
		updatedAt     sql.NullString
	)

	query := `
		SELECT
			book_id,
			book_name,
			author_name,
			book_price,
			created_at,
			updated_at
		FROM
			books
		WHERE book_id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&book_name,
			&author_name,
			&book_price,
			&createdAt,
			&updatedAt,
		)

	if err != nil {
		return nil, err
	}

	return &models.Book{
		Id:           id.String,
		BookName:     book_name.String,
		AuthorName:   author_name.String,
		BookPrice:    book_price.Float64,
		CreatedAt:    createdAt.String,
		UpdatedAt:    updatedAt.String,
	}, nil
}

func (f *BookRepo) GetList(ctx context.Context, req *models.GetListBookRequest) (*models.GetListBookResponse, error) {

	var (
		resp   = models.GetListBookResponse{}
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
			book_id,
			book_name,
			author_name,
			book_price,
			created_at,
			updated_at
		FROM
			books
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {

		var (
			id            sql.NullString
			book_name     sql.NullString
			author_name   sql.NullString
			book_price    sql.NullFloat64
			createdAt     sql.NullString
			updatedAt     sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&book_name,
			&author_name,
			&book_price,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Books = append(resp.Books, &models.Book{
			Id:           id.String,
			BookName:     book_name.String,
			AuthorName:   author_name.String,
			BookPrice:    book_price.Float64,
			CreatedAt:    createdAt.String,
			UpdatedAt:    updatedAt.String,
		})

	}

	return &resp, err
}

func (f *BookRepo) Update(ctx context.Context, req *models.UpdateBook) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			books
		SET
			book_name   = :book_name,
			author_name = :author_name,
			book_price  = :book_price,
			updated_at  = now()
		WHERE book_id   = :book_id
	`

	params = map[string]interface{}{
		"book_id":     req.Id,
		"book_name":   req.BookName,
		"author_name": req.AuthorName,
		"book_price":  req.BookPrice,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *BookRepo) Delete(ctx context.Context, req *models.BookPrimarKey) error {

	_, err := f.db.Exec(ctx, "DELETE FROM books WHERE book_id = $1", req.Id)
	if err != nil {
		return err
	}

	return err
}
