package models

type UserPrimarKey struct {
	Id string `json:"user_id"`
}

type CreateUser struct {
	First_name   string  `json:"first_name"`
	Last_name    string  `json:"last_name"`
	Phone_number string  `json:"phone_number"`
	Balance      float64 `json:"balance"`
}
type User struct {
	Id           string  `json:"user_id"`
	First_name   string  `json:"first_name"`
	Last_name    string  `json:"last_name"`
	Phone_number string  `json:"phone_number"`
	Balance      float64 `json:"balance"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type UpdateUser struct {
	Id           string  `json:"user_id"`
	First_name   string  `json:"first_name"`
	Last_name    string  `json:"last_name"`
	Phone_number string  `json:"phone_number"`
	Balance      float64 `json:"balance"`
}

type GetListUserRequest struct {
	Limit  int32
	Offset int32
}

type GetListUserResponse struct {
	Count int32   `json:"count"`
	Users []*User `json:"users"`
}