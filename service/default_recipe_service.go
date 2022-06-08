package service

import (
	"errors"
	"RecipeCollection/repository"
	"RecipeCollection/model"
)

var (
	ErrIDIsNotValid    = errors.New("Id is not valid")
	ErrTitleIsEmpty = errors.New("Recipe title cannot be empty")
	ErrRecipeNotFound   = errors.New("This recipe cannot be found")
)

type defaultRecipeService struct {
	recipeRepo repository.IRecipeRepository
}

func NewDefaultRecipeService(rRepo repository.IRecipeRepository) *defaultRecipeService {
	return &defaultRecipeService{
		recipeRepo: rRepo,
	}
}

func (d *defaultRecipeService) GetRecipes() ([]model.Recipe, error) {
	return d.recipeRepo.GetRecipes()
}