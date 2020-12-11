package domain

import "time"

type User struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	CreatedDate time.Time  `json:"created_date"`
	DeletedDate *time.Time `json:"deleted_date"`
}

type UserInfo []User

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
