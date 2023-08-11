package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.dev/nicolasmmb/GoExpert-Topicos/configs"
	"github.dev/nicolasmmb/GoExpert-Topicos/internal/entity"
	"github.dev/nicolasmmb/GoExpert-Topicos/internal/infra/database"
	"github.dev/nicolasmmb/GoExpert-Topicos/internal/infra/webserver/handlers"
	"github.dev/nicolasmmb/GoExpert-Topicos/pkg/middlewares"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	env := configs.LoadENVs("./")
	env.Print()

	// DB
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&entity.User{}, &entity.Product{})
	if err != nil {
		log.Fatal(err)
	}

	// gin
	r := chi.NewRouter()
	// Handlers
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler((productDB))

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	// Middlewares
	r.Use(middleware.Logger)
	middlewares.SetJWTExp(env.JWT_EXP)
	middlewares.SetJWTSecretKey(env.JWT_SECRET)

	// -> Routes
	// Products
	r.Route("/products", func(r chi.Router) {

		r.Use(middlewares.JWTVerify)

		r.Post("/", productHandler.Create)
		r.Get("/", productHandler.GetById)
		r.Put("/", productHandler.Update)
		r.Delete("/", productHandler.Delete)
		r.Get("/all", productHandler.GetAll)
	})

	// -> Users
	r.Post("/users/auth", userHandler.GetJWT)
	r.Post("/users", userHandler.Create)
	r.Get("/users", userHandler.GetById)
	r.Put("/users", userHandler.Update)
	r.Delete("/users", userHandler.Delete)
	r.Get("/users/all", userHandler.GetAll)

	// Server
	serverAddress := env.GetServerAddress()
	log.Println("Server running on: " + serverAddress)
	http.ListenAndServe(serverAddress, r)

}
