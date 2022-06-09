package repository

import (
	"RecipeCollection/model"
)

type IRecipeRepository interface {
	GetRecipes() ([]model.Recipe, error)
	GetRecipe(id int) (model.Recipe, error)
	CreateRecipe(recipe model.Recipe) error
	DeleteRecipe(id int) error
}