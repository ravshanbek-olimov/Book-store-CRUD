package models

type OrderPrimarKey struct {
	Id string `json:"order_id"`
}

type CreateOrder struct {
	UserId      string `json:"user_id"`
	BookId      string `json:"book_id"`
	OrderPrice  float64 `json:"order_price"`
}
type Order struct {
	Id         string `json:"order_id"`
	UserId     string `json:"user_id"`
	BookId     string `json:"book_id"`
	OrderPrice float64 `json:"order_price"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type UpdateOrder struct {
	Id          string  `json:"order_id"`
	UserId    	string `json:"user_id"`
	BookId      string `json:"book_id"`
	OrderPrice  float64 `json:"order_price"`
}

type GetListOrderRequest struct {
	Limit  int32
	Offset int32
}

type GetListOrderResponse struct {
	Count  int32    `json:"count"`
	Orders []*Order `json:"orders"`
}


