package models

type BookPrimarKey struct {
	Id string `json:"book_id"`
}

type CreateBook struct {
	BookName   string `json:"book_name"`
	AuthorName string `json:"author_name"`
	BookPrice  float64 `json:"book_price"`
}
type Book struct {
	Id         string `json:"book_id"`
	BookName   string `json:"book_name"`
	AuthorName string `json:"author_name"`
	BookPrice  float64 `json:"book_price"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type UpdateBook struct {
	Id          string  `json:"book_id"`
	BookName 	string  `json:"book_name"`
	AuthorName  string  `json:"author_name"`
	BookPrice   float64 `json:"book_price"`
}

type GetListBookRequest struct {
	Limit  int32
	Offset int32
}

type GetListBookResponse struct {
	Count  int32    `json:"count"`
	Books []*Book `json:"books"`
}


