package route

import (
	"crud/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	ejwt := e.Group("/jwt")
	ejwt.GET("/users", controllers.GetUsersController)
	ejwt.GET("/users/:id", controllers.GetUserController)
	ejwt.POST("/users", controllers.CreateUserController)
	ejwt.POST("/login", controllers.LoginUserController)
	ejwt.DELETE("/users/:id", controllers.DeleteUserController)
	ejwt.PUT("/users/:id", controllers.UpdateUserController)

	ejwt.GET("/books", controllers.GetBooksController)
	ejwt.GET("/books/:id", controllers.GetBookController)
	ejwt.POST("/books", controllers.CreateBookController)
	ejwt.DELETE("/books/:id", controllers.DeleteBookController)
	ejwt.PUT("/books/:id", controllers.UpdateBookController)

	e.GET("/blogs", controllers.GetBlogsController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.POST("/blogs", controllers.CreateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)

	return e
}
