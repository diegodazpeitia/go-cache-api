## 1st, what you can check on the browser

<img width="1436" alt="Captura de pantalla 2024-09-06 a la(s) 12 16 11" src="https://github.com/user-attachments/assets/4c76f199-0439-4b09-829e-02cd106df141">


# Go-Cache-API

## Overview

Go-Cache-API is a simple web application built with Go and the Fiber framework that demonstrates caching functionality using the `go-cache` library. This application includes middleware to cache GET request responses and serve them efficiently. The cache is configured to store responses for 10 minutes, reducing the need for repeated processing of the same requests.

## Features

- **Cache Middleware**: Implements caching for GET requests to improve response time and reduce load on the server.
- **Simple API**: Includes a basic API endpoint to illustrate the caching mechanism.

## Installation

1. **Clone the repository:**

   ```bash
   git clone <repository-url>
   ```

2. **Navigate to the project directory:**

   ```bash
   cd Go-Cache-API
   ```

3. **Install the dependencies:**

   ```bash
   go mod tidy
   ```

## Usage

1. **Run the application:**

   ```bash
   go run main.go
   ```

2. **Access the application:**

   Open your web browser or use a tool like `curl` to access the endpoints:

   - `GET /`: Returns a "Hello, World 👋!" message.
   - `GET /posts/:id`: Demonstrates caching for POST requests. Replace `:id` with an actual post ID.

## Code Explanation

### Middleware

The `CacheMiddleware` function is a Fiber middleware that:

- Checks if the request method is `GET` (caching is only applied to GET requests).
- Generates a cache key based on the request path and query parameters.
- Checks if the response is already cached. If so, it returns the cached response.
- If not cached, it proceeds to handle the request, caches the response, and then returns it.

```go
func CacheMiddleware(cache *cache.Cache) fiber.Handler {
    return func(c *fiber.Ctx) error {
        if c.Method() != "GET" {
            // Only cache GET requests
            return c.Next()
        }

        cacheKey := c.Path() + "?" + c.Params("id") // Generate a cache key from the request path and query parameters

        // Check if the response is already in the cache
        if cached, found := cache.Get(cacheKey); found {
            c.Response().Header.Set("Cache-Status", "HIT")
            return c.JSON(cached)
        }

        c.Set("Cache-Status", "MISS")
        err := c.Next()
        if err != nil {
            return err
        }

        var data Post
        cacheKey = c.Path() + "?" + c.Params("id")

        body := c.Response().Body()
        err = json.Unmarshal(body, &data)
        if err != nil {
            return c.JSON(fiber.Map{"error": err.Error()})
        }

        // Cache the response for 10 minutes
        cache.Set(cacheKey, data, 10*time.Minute)

        return nil
    }
}
```

### Main Application

The main application initializes a Fiber instance and sets up a basic route with caching middleware:

```go
func main() {
    app := fiber.New() // Creating a new instance of Fiber.

    cache := cache.New(10*time.Minute, 20*time.Minute) // setting default expiration time and clearance time.

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })
    app.Get("/posts/:id", middleware.CacheMiddleware(cache), routes.GetPosts) //commenting this route just to test the "/" endpoint.
    app.Listen(":8080")
}
```

## Contributing

Feel free to fork the repository and submit pull requests. For significant changes or improvements, please open an issue to discuss your proposal.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries or further information, please reach out via the repository's issue tracker.
