package main

import (
	"os"

	"github.com/DeepjyotiSarmah/go_MVC_with_React_tutorial/controllers"
	"github.com/DeepjyotiSarmah/go_MVC_with_React_tutorial/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	// Load templates
	engine := html.New("./views", ".html")

	// Create app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Configure app
	app.Static("/", "./public")

	// Routing
	app.Get("/api/tasks", controllers.FetchTasks)
	app.Post("/api/tasks", controllers.CreateTask)
	app.Get("/api/tasks/:id", controllers.FetchTask)
	app.Delete("/api/tasks/:id", controllers.DeleteTask)

	frontendRoutes := []string{
		"/",
		"/about",
	}

	for _, route := range frontendRoutes {
		app.Get(route, controllers.Home)
	}

	app.Get("/", controllers.Home)

	// Start app
	app.Listen(":" + os.Getenv("PORT"))
}
