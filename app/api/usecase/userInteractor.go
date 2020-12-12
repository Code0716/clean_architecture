package usecase

import "github.com/Code0716/clean_architecture/app/api/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (err error) {
	_, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Users() (user domain.UserInfo, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier string) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}
func (interactor *UserInteractor) UserByQuery(Query, param string) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByQuery(Query, param)
	return
}
