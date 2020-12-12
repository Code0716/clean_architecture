package usecase

import "github.com/Code0716/clean_architecture/app/api/domain"

type UserRepository interface {
	Store(domain.User) (string, error)
	FindById(string) (domain.User, error)
	FindByQuery(string, string) (domain.User, error)
	FindAll() (domain.UserInfo, error)
	Delete(domain.User) error
}
