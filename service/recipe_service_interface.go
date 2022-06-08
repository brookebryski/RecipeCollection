package service

import "RecipeCollection/model"

type IRecipeService interface {
	GetRecipes() ([]model.Recipe, error)
}