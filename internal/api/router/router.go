package router

import (
	"fmt"
	"io"
	"os"

	"firebase.google.com/go/auth"
	"github.com/atsur/api-server/internal/api/controllers"
	"github.com/atsur/api-server/internal/api/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	apiKeyName = os.Getenv("API_KEY_NAME")
)

func Setup(client *auth.Client) *gin.Engine {
	app := gin.New()

	// Logging to a file.
	f, _ := os.Create("log/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())

	// Routes
	// ================== Login Routes
	app.POST("/api/login", controllers.Login)
	app.POST("/api/login/add", controllers.CreateUser)
	// ================== Docs Routes
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// FirebaseAuth protected routes
	versioned := app.Group("/v1")
	// Use JWT auth middleware
	versioned.Use(middlewares.AuthJWT(client))
	// ================== Tasks Routes
	versioned.GET("/api/tasks/:id", controllers.GetTaskById)
	versioned.GET("/api/tasks", controllers.GetTasks)
	versioned.POST("/api/tasks", controllers.CreateTask)
	versioned.PUT("/api/tasks/:id", controllers.UpdateTask)
	versioned.DELETE("/api/tasks/:id", controllers.DeleteTask)

	// Admin routes
	admin := app.Group("/v1/admin")
	// must have api key
	admin.Use(middlewares.AuthAPIKey(apiKeyName))
	// ================== User Routes
	admin.GET("/users", controllers.GetUsers)
	admin.GET("/api/users/:id", controllers.GetUserById)
	admin.POST("/api/users", controllers.CreateUser)
	admin.PUT("/api/users/:id", controllers.UpdateUser)
	admin.DELETE("/api/users/:id", controllers.DeleteUser)

	return app
}
