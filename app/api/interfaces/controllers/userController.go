package controllers

import (
	"github.com/Code0716/clean_architecture/app/api/domain"
	"github.com/Code0716/clean_architecture/app/api/interfaces/database"
	"github.com/Code0716/clean_architecture/app/api/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context, uuid string, time time.Time) {
	u := new(domain.User)
	c.Bind(&u)
	u.ID = uuid
	u.CreatedDate = time
	err := controller.Interactor.Add(*u)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, err)
}

func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, users)
}

func (controller *UserController) Show(c Context) {
	id := c.Param("id")
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
}
