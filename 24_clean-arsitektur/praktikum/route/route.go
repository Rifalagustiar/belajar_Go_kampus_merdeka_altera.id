package route

import (
	"crud/controller"
	"crud/repository"
	"crud/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

const SECRET_JWT = "legal"

func NewRoute(e *echo.Echo, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository)

	userController := controller.NewUserController(userUsecase)

	e.POST("/users", userController.Create)
	e.POST("/login", userController.LoginUser)

	r := e.Group("/jwt")
	{
		jwtConfig := middleware.JWTConfig{
			SigningKey: []byte(SECRET_JWT),
		}
		r.Use(middleware.JWTWithConfig(jwtConfig))
		r.GET("/users", userController.GetAllUsers)

	}
}
