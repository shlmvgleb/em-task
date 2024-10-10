package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shlmvgleb/em-task/cmd/docs"
	"github.com/shlmvgleb/em-task/internal/config"
	"github.com/shlmvgleb/em-task/internal/database"
	"github.com/shlmvgleb/em-task/internal/handlers"
	repositories "github.com/shlmvgleb/em-task/internal/repositories/postgres"
	"github.com/shlmvgleb/em-task/internal/services"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	DevEnv  = "development"
	ProdEnv = "production"
)

func main() {
	config := config.ReadFromEnv()
	loggerSetup(config)

	ctx := context.Background()
	db, err := database.New(config.Postgres, ctx)
	if err != nil {
		log.Fatalf("error while connecting to database: %s", err)
	}

	songRepo := repositories.NewPostgresSongRepo(db)

	songService := services.NewSongService(songRepo)
	songDetailsApiService := services.NewSongDetailsMockApiService()

	cntrl := handlers.NewController(
		songService,
		songDetailsApiService,
	)

	err = startServer(config, cntrl)
	if err != nil {
		log.Fatalf("server is abruptly closed: %s", err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func loggerSetup(config *config.AppConfig) {
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint:      true,
		DisableTimestamp: false,
	})

	if config.AppEnv == DevEnv {
		log.SetLevel(log.DebugLevel)
	}

	if config.AppEnv == ProdEnv {
		log.SetLevel(log.InfoLevel)
	}
}

func startServer(config *config.AppConfig, cntrl *handlers.Controller) error {
	if config.AppEnv == DevEnv {
		gin.SetMode("debug")
	}

	if config.AppEnv == ProdEnv {
		gin.SetMode("release")
	}

	engine := gin.Default()
	engine.Use(CORSMiddleware())

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := engine.Group("/api/v1")
	{
		sg := v1.Group("/songs")
		{
			sg.POST("/", cntrl.AddSong)
			sg.PATCH("/", cntrl.UpdateSong)
			sg.DELETE("/:id", cntrl.DeleteSong)
			sg.GET("/", cntrl.GetSongsWithPagination)
			sg.GET("/:id", cntrl.GetSongByIdWithVersePagination)
		}
	}

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Infof("Server listening on port: %d", config.Port)
	return engine.Run(fmt.Sprintf(":%v", config.Port))
}
