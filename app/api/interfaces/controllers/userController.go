package controllers

import (
	"time"

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

func (controller *UserController) Create(c Context, uuid string, createTime time.Time, hashFunc func(string) (string, error), getNewToken func(string, string, string) string) {
	u := new(domain.User)
	c.Bind(&u)
	u.ID = uuid
	hashedPass, _ := hashFunc(u.Password)
	u.Password = hashedPass
	u.CreatedDate = createTime
	err := controller.Interactor.Add(*u)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	tokenString := getNewToken(u.ID, u.Name, u.Email)
	response := make(map[string]string)
	response["authorization"] = tokenString
	c.JSON(200, response)
}

func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	response := make(map[string]domain.UserInfo, len(users))
	response["data"] = users
	c.JSON(200, response)
}

func (controller *UserController) Show(c Context) {
	id := c.Param("id")
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, user)
}
