package repository

import (
	"errors"
	"RecipeCollection/model"
)

var (
	ErrRecipeNotFound = errors.New("From repository - recipe not found.")
)

type inMemoryRecipeRepository struct {
	Recipes []model.Recipe
}

func NewInMemoryRecipeRepository() *inMemoryRecipeRepository {
	var recipes = []model.Recipe {
		{ID: 1, Title: "Soft and Chewy Chocolate Chip Cookies", Author: "Martha Stewart", Rating: 4},
		{ID: 2, Title: "No-Bake Cheesecake", Author: "Martha Stewart", Rating: 4},
		{ID: 3, Title: "Peas and Pancetta", Author: "Ina Garten", Rating: 4},
		{ID: 4, Title: "Truffled Mac and Cheese", Author: "Ina Garten", Rating: 4},
	}
	return &inMemoryRecipeRepository{
		Recipes: recipes,
	}
}

func(i *inMemoryRecipeRepository) GetRecipes() ([]model.Recipe, error) {
	return i.Recipes, nil
}
	