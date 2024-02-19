# Backend
## Getting Started
1. [Install Go](https://go.dev/doc/install)
2. [Install Docker](https://docs.docker.com/get-docker/)
   - Run mongoDB container
        ```bash
         docker run --name mongo -d -p 27017:27017 mongodb/mongodb-community-server:latest
        ```
3. Install dependencies
    ```bash
    go mod get
    ```
   
4. Copy the `.env.example` file to `.env` and fill in the fields to match your setup
    ```bash
    # For Mac/Linux
    cp .env.example .env
   
    # For Windows
    copy .env.example .env
    ```
5. Run the server
    ```bash
    go run main.go
    ```

## Dependencies Used
- [Gin](https://gin-gonic.com/docs/introduction/)
  - [GoDoc](https://pkg.go.dev/github.com/gin-gonic/gin)
- [MongoDB Go Driver](https://www.mongodb.com/docs/drivers/go/current/)
  - [GoDoc](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo)
- [JWT Go](https://pkg.go.dev/github.com/golang-jwt/jwt/v4)

## Useful Tools
- API/REST Client
  - [Thunder Client for VS Code](https://www.thunderclient.io/)
  - [GoLand HTTP Client](https://www.jetbrains.com/help/idea/http-client-in-product-code-editor.html)
  - [Insomnia](https://insomnia.rest/download)
  - [Postman](https://www.postman.com/downloads/)
- Database Client
  - [MongoDB Compass](https://www.mongodb.com/try/download/compass)
  - [MongoDB for VS Code](https://marketplace.visualstudio.com/items?itemName=mongodb.mongodb-vscode)
  - [GoLand Database Tools](https://www.jetbrains.com/help/idea/mongodb.html)

## Project Structure
### When to use what (TL;DR)
- **Routes**: Use routes to define the endpoints needed to connect the frontend
- **Models**: Use models to define the data that needs to e modeled
  - Think of these as the data structures that will be used in the application
- **Controllers**: Use controllers to define the logic that will be used to interact with the models
  - Typically, this is where most database interactions will occur. Data will also typically be formatted here before being sent to the frontend
- **Middleware**: Use middleware to define any logic that needs to be run before or after a request is handled
  - This can include things like authentication, logging, or error handling
### Routes
Routes should be defined in the `routes` package. Each route should be defined in its own file and should be grouped by the type of route (e.g. `auth.go`, `user.go`, `post.go`).
#### Grouping Routes
Routes should be grouped by using the `Group` method on the router. This is useful for grouping routes that share a common prefix or middleware.
```go
// Create a new router group for "/api" routes
posts := router.Group("/posts")

// Define a GET route for "/home" in the "/api" group
posts.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{"data": "List of posts"})
})
```
#### Authenticated Routes
Routes that require authentication should also be defined in a group. This group should be protected by the `Authorized` middleware.
```go
// Create a new router group for authenticated routes
profile := router.Group("/profile")
profile.Use(middleware.Authorized())

// Define a GET route for "/home" in the authenticated group
profile.GET("/me", func(c *gin.Context) {
    c.JSON(200, gin.H{"data": "User profile"})
})
```
### Models
Models should be defined in the `models` package. Each model should be defined in its own file and should be grouped by the type of model (e.g. `user.go`, `post.go`).
### Controllers
Controllers should be defined in the `controllers` package. Each controller should be defined in its own file and should be grouped by the type of controller (e.g. `auth.go`, `user.go`, `post.go`).
#### Controller Methods
Each controller method should be defined as a function that takes a `*gin.Context` as a parameter. This allows the method to be used as a route handler.
```go
// Define a controller method for the "/home" route
func Home(c *gin.Context) {
    return c.JSON(200, gin.H{"data": "Welcome home!"})
}
```

### Middleware
Middleware should be defined in the `middleware` package. Each middleware should be defined in its own file and should be grouped by the type of middleware (e.g. `auth.go`, `logging.go`, `error.go`).