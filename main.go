package main

import (
	"goNew/database"
	"goNew/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)
	app.Get("/fact", handlers.NewFactView)
	app.Post("/fact", handlers.CreateFact)
}

func main() {
	database.ConnectDb()

	engine := html.New("./views", ".html") // 2. new engine

	app := fiber.New(fiber.Config{
		Views:       engine,         // new config
		ViewsLayout: "layouts/main", // add this to config
	})

	setupRoutes(app)

	app.Listen(":3000")

}
