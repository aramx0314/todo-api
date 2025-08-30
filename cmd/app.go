package main

import (
	"log"
	"net/http"
	"todo-api/internal/db"
	"todo-api/internal/handler"
	"todo-api/internal/middleware"
	"todo-api/internal/repository"
	"todo-api/internal/service"

	_ "todo-api/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title TODO-API
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT Authorization: "Bearer {token}"
func main() {
	db, err := db.Init()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	r.Use(middleware.JsonResponseHeaderMiddleware)

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)
	authHandler.RegisterRoutes(r)

	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)
	todoHandler.RegisterRoutes(r)

	log.Println("start server: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
