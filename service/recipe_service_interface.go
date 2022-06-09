package service

import "RecipeCollection/model"

type IRecipeService interface {
	GetRecipes() ([]model.Recipe, error)
	GetRecipe(id int) (model.Recipe, error)
	CreateRecipe(recipe model.Recipe) error
	DeleteRecipe(id int) error
}