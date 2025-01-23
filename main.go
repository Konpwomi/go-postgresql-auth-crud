package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/konpwomi/go-postgresql-auth-crud/api/routes"
)

func main() {
    app := fiber.New()

    // Register routes
    routes.RegisterRoutes(app)

    // Start server
    log.Fatal(app.Listen(":8080"))
}
