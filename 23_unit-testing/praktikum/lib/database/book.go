package database

import (
	"crud/config"
	"crud/models"

	"github.com/labstack/echo/v4"
)

func GetBooks() (interface{}, error) {
	var books []models.Books

	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func CreateBook(c echo.Context) (interface{}, error) {
	var book models.Books
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func GetBook(bookID int) (interface{}, error) {
	// query the database for the book with the given ID
	var book models.Books

	if err := config.DB.Where("id = ?", bookID).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func DeleteBook(bookID int) (interface{}, error) {
	var book models.Books
	// delete the book with the given ID from the database

	if err := config.DB.Where("id = ?", bookID).Delete(&book).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

func UpdateBook(c echo.Context) (interface{}, error) {
	var book models.Books

	if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		return nil, err
	}
	if err := c.Bind(&book); err != nil {
		return nil, err
	}
	config.DB.Save(&book)

	return book, nil
}
