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

	// ================== Entry Routes
	versioned.GET("/api/entry/:id", controllers.GetEntryById)
	versioned.GET("/api/entries", controllers.GetEntries)
	versioned.POST("/api/entry", controllers.CreateEntry)
	versioned.PUT("/api/entry/:id", controllers.UpdateEntry)
	versioned.DELETE("/api/entry/:id", controllers.DeleteEntry)

	// ================== Record Routes
	versioned.GET("/api/record/:id", controllers.GetRecordById)
	versioned.GET("/api/records", controllers.GetRecords)
	versioned.POST("/api/record", controllers.CreateRecord)
	versioned.PUT("/api/record/:id", controllers.UpdateRecord)
	versioned.DELETE("/api/record/:id", controllers.DeleteRecord)

	// ================== Record Routes
	versioned.GET("/api/record/:id/provenance", controllers.GetProvenanceByRecordId)
	versioned.GET("/api/record/:id/exhibitions", controllers.GetExhibitionByRecordId)
	versioned.GET("/api/record/:id/auctions", controllers.GetAuctionsByRecordId)

	// ================== Provenance Routes
	versioned.GET("/api/provenance/:id", controllers.GetProvenanceById)
	versioned.GET("/api/provenance", controllers.GetProvenances)
	versioned.POST("/api/provenance", controllers.CreateProvenance)
	versioned.PUT("/api/provenance/:id", controllers.UpdateProvenance)
	versioned.DELETE("/api/provenance/:id", controllers.DeleteProvenance)

	// ================== Exhibition Routes
	versioned.GET("/api/exhibition/:id", controllers.GetExhibitionById)
	versioned.GET("/api/exhibition", controllers.GetExhibitions)
	versioned.POST("/api/exhibition", controllers.CreateExhibition)
	versioned.PUT("/api/exhibition/:id", controllers.UpdateExhibition)
	versioned.DELETE("/api/exhibition/:id", controllers.DeleteExhibition)

	// ================== Auction Routes
	versioned.GET("/api/auction/:id", controllers.GetAuctionById)
	versioned.GET("/api/auction", controllers.GetAuctions)
	versioned.POST("/api/auction", controllers.CreateAuction)
	versioned.PUT("/api/auction/:id", controllers.UpdateAuction)
	versioned.DELETE("/api/auction/:id", controllers.DeleteAuction)

	// ================== Profile Routes
	versioned.GET("/api/profile/:id", controllers.GetProfileById)
	versioned.GET("/api/profile/user/:user_id", controllers.GetProfileByUserId)
	versioned.GET("/api/profile", controllers.GetProfiles)
	versioned.POST("/api/profile", controllers.CreateProfile)
	versioned.PUT("/api/profile/:id", controllers.UpdateProfile)
	versioned.DELETE("/api/profile/:id", controllers.DeleteProfile)

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

	admin.POST("/api/qr-code", controllers.CreateQR)

	// ================== Record Routes
	admin.GET("/api/Record/:id", controllers.GetRecordById)
	admin.GET("/api/Records", controllers.GetRecords)
	admin.POST("/api/Record", controllers.CreateRecord)
	admin.PUT("/api/Record/:id", controllers.UpdateRecord)
	admin.DELETE("/api/Record/:id", controllers.DeleteRecord)

	return app
}
