package main

import (
	"log"
	"time"

	"github.com/arashiaslan/forum-app-go/internal/configs"
	"github.com/arashiaslan/forum-app-go/internal/handlers/memberships"
	"github.com/arashiaslan/forum-app-go/internal/handlers/posts"
	"github.com/arashiaslan/forum-app-go/pkg/internalsql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	membershipRepo "github.com/arashiaslan/forum-app-go/internal/repository/memberships"
	postRepo "github.com/arashiaslan/forum-app-go/internal/repository/posts"
	membershipSvc "github.com/arashiaslan/forum-app-go/internal/service/memberships"
	postSvc "github.com/arashiaslan/forum-app-go/internal/service/posts"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal membaca konfigurasi", err)
	}

	cfg = configs.Get()
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// CORS middleware - menyesuaikan origin dev Vite/React
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"}, // sesuaikan origin frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// logging & recovery
	r.Use(gin.Logger(), gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipsHandler := memberships.NewHandler(r, membershipService)
	membershipsHandler.RegisterRoute()

	postsHandler := posts.NewHandler(r, postService)
	postsHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}