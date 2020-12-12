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
	response["Authorization"] = tokenString
	c.JSON(200, response)
}

func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	var userData []domain.UserResponse
	for _, value := range users {
		user := new(domain.UserResponse)
		user.ID = value.ID
		user.Name = value.Name
		user.Email = value.Email
		userData = append(userData, *user)
	}
	response := make(map[string][]domain.UserResponse)

	response["data"] = userData
	c.JSON(200, response)
}

func (controller *UserController) Show(c Context) {
	id := c.Param("id")
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	userResponse := new(domain.UserResponse)
	userResponse.ID = user.ID

	userResponse.Name = user.Name
	userResponse.Email = user.Email

	response := make(map[string]domain.UserResponse)
	response["data"] = *userResponse
	c.JSON(200, response)
}

func (controller *UserController) Login(c Context, passwordVerify func(hash, pw string) error, getNewToken func(string, string, string) string) {

	u := new(domain.User)
	c.Bind(&u)
	response := make(map[string]string)

	// 定数化する？
	user, err := controller.Interactor.UserByQuery("Email", u.Email)
	response["error"] = "Not match email or password"
	if err != nil {
		c.JSON(500, response)
		return
	}
	err = passwordVerify(user.Password, u.Password)
	if err != nil {
		c.JSON(500, response)
		return
	}
	tokenString := getNewToken(user.ID, user.Name, user.Email)
	response["Authorization"] = tokenString
	response["error"] = ""
	c.JSON(200, response)
}
