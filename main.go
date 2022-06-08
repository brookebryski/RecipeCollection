package main

import (
	"RecipeCollection/handler"
	"RecipeCollection/repository"
	"RecipeCollection/service"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"	
)

func main() {
	recipeInMemoryRepository := repository.NewInMemoryRecipeRepository()
	recipeService := service.NewDefaultRecipeService(recipeInMemoryRepository)
	recipeHandler := handler.NewRecipeHandler(recipeService)

	router := httprouter.New()

	router.GET("/recipes", recipeHandler.GetRecipes)

	log.Println("http server runs on :8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}

// Handler: Layer that gets http request and returns http response to the client.
// Service: Layer that our business logic is in.
// Repository: Layer that provides all necessary data from external (DBs) or internal (in-memory) data source.