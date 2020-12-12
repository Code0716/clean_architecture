package controllers

import (
	"net/http"
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
	response := AuthResponse{}
	response.Status = http.StatusInternalServerError
	_, err := controller.Interactor.UserByQuery("Email", u.Email)
	if err == nil {
		response.ErrorMessage = "User already exist"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	u.ID = uuid
	hashedPass, _ := hashFunc(u.Password)
	u.Password = hashedPass
	u.CreatedDate = createTime
	err = controller.Interactor.Add(*u)
	if err != nil {
		response.ErrorMessage = "Fatal error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	tokenString := getNewToken(u.ID, u.Name, u.Email)
	response.Authorization = tokenString
	response.ErrorMessage = ""
	response.Status = http.StatusOK
	c.JSON(http.StatusOK, response)
}

func (controller *UserController) Index(c Context) {
	response := Response{}
	response.Status = http.StatusInternalServerError
	users, err := controller.Interactor.Users()

	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
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

	response.Data = &userData
	response.Status = http.StatusOK
	c.JSON(http.StatusOK, response)
}

func (controller *UserController) Show(c Context) {
	response := Response{}
	response.Status = http.StatusInternalServerError
	id := c.Param("id")
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	userResponse := new(domain.UserResponse)
	userResponse.ID = user.ID

	userResponse.Name = user.Name
	userResponse.Email = user.Email

	response.Data = &userResponse
	response.Status = http.StatusOK
	c.JSON(http.StatusOK, response)
}

func (controller *UserController) Login(c Context, passwordVerify func(hash, pw string) error, getNewToken func(string, string, string) string) {
	response := AuthResponse{}
	response.Status = http.StatusInternalServerError

	u := new(domain.User)
	c.Bind(&u)

	// 定数化する？
	user, err := controller.Interactor.UserByQuery("Email", u.Email)
	response.ErrorMessage = "Not match email or password"
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = passwordVerify(user.Password, u.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	tokenString := getNewToken(user.ID, user.Name, user.Email)
	response.Authorization = tokenString
	response.ErrorMessage = ""
	response.Status = http.StatusOK
	c.JSON(http.StatusOK, response)
}
