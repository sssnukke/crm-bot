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
	employeeRepo := repository.NewEmployeeRepository(db)

	userService := service.NewUserService(userRepo)
	employeeService := service.NewEmployeeService(employeeRepo, userRepo)

	userHandler := handlers.NewUserHandler(userService)
	employeeHandler := handlers.NewEmployeeHandler(employeeService)

	protected := http.NewServeMux()

	protected.HandleFunc("POST /users", userHandler.CreateUser)
	protected.HandleFunc("GET /users/check", userHandler.CheckUserExists)
	protected.HandleFunc("GET /users", userHandler.GetUserByTgID)

	protected.HandleFunc("POST /employees", employeeHandler.CreateEmployee)

	mux.Handle("/", middleware.AuthMiddleware(cfg.SecretKey, protected))

	return mux
}
