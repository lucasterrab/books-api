package main

import (
	"fmt"

	"github.com/lucasterrab/books-api/book"
	"github.com/lucasterrab/books-api/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"

	// sqlite support
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "book.db")
	if err != nil {
		panic("failed to connect dabase")
	}
	fmt.Println("Database connection succesfully opened")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)
}