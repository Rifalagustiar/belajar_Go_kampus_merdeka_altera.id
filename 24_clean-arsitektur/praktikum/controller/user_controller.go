package controller

import (
	"crud/middleware"
	"crud/models"
	"crud/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController interface{}

type userController struct {
	useCase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *userController {
	return &userController{useCase: userUsecase}
}

func (u *userController) Create(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := u.useCase.Create(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Success Create User",
	})
}

func (u *userController) GetAllUsers(c echo.Context) error {
	userId := middleware.ExtractTokenUserId(c)
	if userId == 0 {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"status":  http.StatusUnauthorized,
			"message": "Token Unauthorized",
		})
	}

	users, err := u.useCase.GetAllUsers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)

	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success Get All User",
		"data":    users,
	})

}

func (u *userController) LoginUser(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	users, err := u.useCase.LoginUser(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	createToken, err := middleware.CreateToken(int(users.ID))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Success Create User",
		"token":   createToken,
		"data":    users,
	})
}
