package main

import (
	"os"

	"github.com/DeepjyotiSarmah/go_MVC_with_React_tutorial/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func main() {
	//Load templates
	engine := html.New("./views", ".html")

	//Setup app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//Config App
	app.Static("/", "./public")

	//Routes
	Routes(app)

	//Start app
	app.Listen(":" + os.Getenv("PORT"))
}
