package repository

import (
	"RecipeCollection/model"
)

type IRecipeRepository interface {
	GetRecipes() ([]model.Recipe, error)
}