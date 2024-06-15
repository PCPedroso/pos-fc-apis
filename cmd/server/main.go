package main

import (
	"net/http"

	"github.com/PCPedroso/pos-fc-apis/configs"
	"github.com/PCPedroso/pos-fc-apis/internal/entity"
	"github.com/PCPedroso/pos-fc-apis/internal/infra/database"
	"github.com/PCPedroso/pos-fc-apis/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, config.TokenAuth, config.JwtExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.Create)
	r.Get("/products", productHandler.FindAll)
	r.Get("/products/{id}", productHandler.FindByID)
	r.Put("/products/{id}", productHandler.Update)
	r.Delete("/products/{id}", productHandler.Delete)

	r.Post("/users", userHandler.Create)
	r.Post("/users/gen_token", userHandler.GetJwt)

	http.ListenAndServe(":8080", r)
}
