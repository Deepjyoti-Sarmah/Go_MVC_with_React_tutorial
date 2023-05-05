package main

import (
	"github.com/DeepjyotiSarmah/go_MVC_with_React_tutorial/controllers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.PostIndex)
}
