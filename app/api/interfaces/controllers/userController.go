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
	Base
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

	_, err := controller.Interactor.UserByQuery("Email", u.Email)
	if err == nil {
		response := controller.Base.FormatResponse(http.StatusInternalServerError)
		response.Meta.ErrorMessage = "User already exist"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	u.ID = uuid
	hashedPass, _ := hashFunc(u.Password)
	u.Password = hashedPass
	u.CreatedDate = createTime
	err = controller.Interactor.Add(*u)
	if err != nil {
		response := controller.Base.FormatResponse(http.StatusInternalServerError)
		response.Meta.ErrorMessage = "Fatal error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	tokenString := getNewToken(u.ID, u.Name, u.Email)
	response := controller.Base.FormatAuthResponse(http.StatusOK, tokenString)
	c.JSON(http.StatusOK, response)
}

func (controller *UserController) Index(c Context) {

	users, err := controller.Interactor.Users()

	if err != nil {
		response := controller.Base.FormatResponse(http.StatusInternalServerError)
		response.Meta.ErrorMessage = "Unkown error"
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
	response := controller.Base.FormatResponse(http.StatusOK, userData)

	c.JSON(http.StatusOK, response)
}

func (controller *UserController) Show(c Context) {

	id := c.Param("id")
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		response := controller.Base.FormatResponse(http.StatusInternalServerError)
		response.Meta.ErrorMessage = "User not found"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	userResponse := new(domain.UserResponse)
	userResponse.ID = user.ID

	userResponse.Name = user.Name
	userResponse.Email = user.Email
	response := controller.Base.FormatResponse(http.StatusOK, userResponse)
	c.JSON(http.StatusOK, response)
}

func (controller *UserController) Login(c Context, passwordVerify func(hash, pw string) error, getNewToken func(string, string, string) string) {

	u := new(domain.User)
	c.Bind(&u)

	// 定数化する？
	user, err := controller.Interactor.UserByQuery("Email", u.Email)
	errorMessage := "Not match email or password"
	if err != nil {
		response := controller.Base.FormatResponse(http.StatusInternalServerError)
		response.Meta.ErrorMessage = errorMessage
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = passwordVerify(user.Password, u.Password)
	if err != nil {
		response := controller.Base.FormatResponse(http.StatusInternalServerError)
		response.Meta.ErrorMessage = errorMessage
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	tokenString := getNewToken(user.ID, user.Name, user.Email)
	response := controller.Base.FormatAuthResponse(http.StatusOK, tokenString)
	c.JSON(http.StatusOK, response)
}

func (controller *UserController) LogicalDelete(c Context, DeleteTime time.Time) {
	id := c.Param("id")
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		response := controller.Base.FormatResponse(http.StatusInternalServerError)
		response.Meta.ErrorMessage = "User not exist"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	user.DeletedDate = &DeleteTime
	err = controller.Interactor.Delete(user)
	if err != nil {
		response := controller.Base.FormatResponse(http.StatusInternalServerError)
		response.Meta.ErrorMessage = "Unkown error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := controller.Base.FormatResponse(http.StatusOK, user)
	c.JSON(http.StatusOK, response)
}
