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
	"gorm.io/gorm/logger"
)

func main() {
	env := configs.LoadENVs("./.env")
	env.Print()

	log.SetPrefix("[MAIN] ")
	log.Println("Starting Database")

	// DB
	logType := env.GetLoggerType()
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logType),
	})
	if err != nil {
		panic(err)
	}

	// Migrations
	log.Println("Running Migrations")
	err = db.AutoMigrate(&entity.User{}, &entity.Product{}, &entity.Vehicle{})
	if err != nil {
		log.Fatal(err)
	}
	configs.CreateAdmin(db)

	// populate.LoadCSV(db)
	log.Println("Starting Server")

	// GIN
	r := chi.NewRouter()
	// Handlers
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler((productDB))

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	vehicleDB := database.NewVehicle(db)
	vehicleHandler := handlers.NewVehicleHandler(vehicleDB)

	// Middlewares
	r.Use(middlewares.CORS)
	r.Use(middlewares.RequestId)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)

	middlewares.SetJWTExp(env.JWT_EXP)
	middlewares.SetJWTSecretKey(env.JWT_SECRET)

	// Routes
	// -> Products
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
	r.Route("/users", func(r chi.Router) {
		r.Use(middlewares.JWTVerify)

		r.Post("/", userHandler.Create)
		r.Get("/", userHandler.GetById)
		r.Put("/", userHandler.Update)
		r.Delete("/", userHandler.Delete)
		r.Get("/all", userHandler.GetAll)
	})

	// -> Vehicles
	r.Route("/vehicles", func(r chi.Router) {
		r.Get("/all", vehicleHandler.FindAll)
	})

	// Server
	serverAddress := env.GetServerAddress()
	log.Println("Server Running: " + serverAddress)
	configs.PrintSeparator()

	http.ListenAndServe(serverAddress, r)

}
