package http

import (
	config "go-wallet/internal/shared/src/infrastructure/configuration"
	"go-wallet/internal/shared/src/infrastructure/database"
	"go-wallet/internal/shared/src/interfaces/http/dependencies"
	"go-wallet/internal/shared/src/interfaces/http/router"
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BuildApp() *gin.Engine {
	_ = godotenv.Load()

	config, err := config.Load()

	if err != nil {
		log.Fatalf("Can't load configuration: %s", err)
	}

	db, err := database.Connect(*config)

	if err != nil {
		log.Fatalf("Can't connect to database: %s", err)
	}

	logger := slog.Default()
	dependencies := dependencies.NewDependencies(db, logger, config)

	r := router.NewRouter(dependencies)

	return r
}
