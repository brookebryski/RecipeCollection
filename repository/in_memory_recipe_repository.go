package repository

import (
	"errors"
	"RecipeCollection/model"
)

var (
	ErrRecipeNotFound = errors.New("from repository - recipe not found")
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
	
func(i *inMemoryRecipeRepository) GetRecipe(id int) (model.Recipe, error) {
	for _, recipe := range i.Recipes {
		if recipe.ID == id {
			return recipe, nil
		}
	}
	return model.Recipe{}, ErrRecipeNotFound
}

func (i *inMemoryRecipeRepository) CreateRecipe(recipe model.Recipe) error {
	recipe.ID = len(i.Recipes) + 1
	i.Recipes = append(i.Recipes, recipe)

	return nil
}

func (i *inMemoryRecipeRepository) DeleteRecipe(id int) error {
	recipeExist := false

	var newRecipeList []model.Recipe
	for _, recipe := range i.Recipes {
		if recipe.ID == id {
			recipeExist = true
		} else {
			newRecipeList = append(newRecipeList, recipe)
		}
	}

	if !recipeExist {
		return ErrRecipeNotFound
	}

	i.Recipes = newRecipeList

	return nil
}