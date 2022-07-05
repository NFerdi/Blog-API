package src

import (
	"blog-api/src/delivery"
	"blog-api/src/middleware"
	"blog-api/src/repository"
	"blog-api/src/usecase"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type server struct {
	db     *gorm.DB
	router *mux.Router
}

type Server interface {
	Run()
}

func InitServer(db *gorm.DB) *server {
	r := mux.NewRouter().StrictSlash(true)

	r.Use(middleware.Logging)

	return &server{
		db:     db,
		router: r,
	}
}

func (c *server) Run() {
	authRequired := middleware.NewAuthMiddleware(c.db)

	userRepository := repository.NewUserRepository(c.db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userDelivery := delivery.NewUserDelivery(userUsecase)
	userDelivery.Mount(c.router, authRequired)

	categoryRepository := repository.NewCategoryRepository(c.db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)
	categoryDelivery := delivery.NewCategoryDelivery(categoryUsecase)
	categoryDelivery.Mount(c.router, authRequired)

	postRepository := repository.NewPostRepository(c.db)
	postUsecase := usecase.NewPostUsecase(postRepository, categoryRepository)
	postDelivery := delivery.NewPostDelivery(postUsecase)
	postDelivery.Mount(c.router, authRequired)

	log.Printf("Server running on port %s", os.Getenv("APP_PORT"))

	if err := http.ListenAndServe(fmt.Sprintf("localhost:%s", os.Getenv("APP_PORT")), c.router); err == nil {
		log.Fatal(err.Error())
	}
}
