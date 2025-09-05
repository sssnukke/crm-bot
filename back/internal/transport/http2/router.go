package http2

import (
	"back/internal/config"
	"back/internal/middleware"
	"back/internal/repository"
	"back/internal/service"
	"back/internal/transport/http2/handlers"
	"net/http"

	"gorm.io/gorm"
)

func NewRouter(cfg *config.Config, db *gorm.DB) *http.ServeMux {
	mux := http.NewServeMux()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	protected := http.NewServeMux()
	protected.HandleFunc("/users", userHandler.CreateUser)

	mux.Handle("/users", middleware.AuthMiddleware(cfg.SecretKey, protected))

	return mux
}
