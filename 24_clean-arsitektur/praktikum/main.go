package main

import (
	"crud/models"
	"crud/route"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/crud_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		models.User{},
	)
	if err != nil {
		panic(err)
	}

	app := echo.New()

	route.NewRoute(app, db)

	app.Logger.Fatal(app.Start(":8080"))
}
