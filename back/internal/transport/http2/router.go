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
	taskRepo := repository.NewTaskRepository(db)

	userService := service.NewUserService(userRepo)
	employeeService := service.NewEmployeeService(employeeRepo, userRepo, "./uploads")
	taskService := service.NewTaskService(taskRepo, employeeRepo)

	userHandler := handlers.NewUserHandler(userService)
	employeeHandler := handlers.NewEmployeeHandler(employeeService)
	taskHandler := handlers.NewTaskHandler(taskService)

	protected := http.NewServeMux()

	protected.HandleFunc("GET /users/check", userHandler.CheckUserExists)
	protected.HandleFunc("POST /users", userHandler.CreateUser)
	protected.HandleFunc("GET /users", userHandler.GetUserByTgID)

	protected.HandleFunc("POST /employees", employeeHandler.CreateEmployee)
	protected.HandleFunc("DELETE /employees", employeeHandler.DeleteEmployeeById)
	protected.HandleFunc("GET /employees", employeeHandler.GetEmployeeById)
	protected.HandleFunc("PATCH /employees", employeeHandler.UpdateEmployeePartial)

	protected.HandleFunc("GET /tasks/employee", taskHandler.GetTasksByEmployeeID)
	protected.HandleFunc("POST /tasks", taskHandler.CreateTask)
	protected.HandleFunc("GET /tasks", taskHandler.GetTaskByID)
	protected.HandleFunc("DELETE /tasks", taskHandler.DeleteTaskByID)

	mux.Handle("/", middleware.AuthMiddleware(cfg.SecretKey, protected))

	fs := http.FileServer(http.Dir("./uploads"))
	protectedFS := middleware.AuthMiddleware(cfg.SecretKey, fs)
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", protectedFS))

	return mux
}
