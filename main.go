package main

import (
	"github.com/diegodazpeitia/Go-Cache-API/middleware"
	"github.com/diegodazpeitia/Go-Cache-API/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	app := fiber.New() // Creating a new instance of Fiber.

	cache := cache.New(10*time.Minute, 20*time.Minute) // setting default expiration time and clearance time.

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/posts/:id", middleware.CacheMiddleware(cache), routes.GetPosts) //commenting this route just to test the "/" endpoint.
	app.Listen(":8080")
}
