package main

import (
	"log"

	"github.com/arashiaslan/forum-app-go/internal/configs"
	"github.com/arashiaslan/forum-app-go/internal/handlers/memberships"
	"github.com/arashiaslan/forum-app-go/pkg/internalsql"
	"github.com/gin-gonic/gin"

	membershipRepo "github.com/arashiaslan/forum-app-go/internal/repository/memberships"
	membershipSvc "github.com/arashiaslan/forum-app-go/internal/service/memberships"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Configs
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

	db, err:= internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	membershipRepo := membershipRepo.NewRepository(db)
	membershipService := membershipSvc.NewService(membershipRepo)

	membershipsHandler := memberships.NewHandler(r, membershipService)
	membershipsHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
