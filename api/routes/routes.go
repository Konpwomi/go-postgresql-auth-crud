package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/konpwomi/go-postgresql-auth-crud/api/handlers"
)

func RegisterRoutes(app *fiber.App) {
    app.Get("/Tasks", handlers.GetTasks)
    app.Get("/Task/:id", handlers.GetTask)
    app.Post("/Task", handlers.CreateTask)
    app.Put("/Task/:id", handlers.UpdateTask)
    app.Delete("/Task/:id", handlers.DeleteTask)
}
