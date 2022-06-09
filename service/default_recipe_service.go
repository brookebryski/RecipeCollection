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

type DefaultRecipeService struct {
	recipeRepo repository.IRecipeRepository
}

func NewDefaultRecipeService(rRepo repository.IRecipeRepository) *DefaultRecipeService {
	return &DefaultRecipeService{
		recipeRepo: rRepo,
	}
}

func (d *DefaultRecipeService) GetRecipes() ([]model.Recipe, error) {
	return d.recipeRepo.GetRecipes()
}

func (d *DefaultRecipeService) GetRecipe(id int) (model.Recipe, error) {
	if id <= 0 {
		return model.Recipe{}, ErrIDIsNotValid
	}
	recipe, err := d.recipeRepo.GetRecipe(id)

	if err != nil {
		if errors.Is(err, repository.ErrRecipeNotFound) {
			return model.Recipe{}, ErrRecipeNotFound
		}
	}
	return recipe, nil
}

func (d *DefaultRecipeService) CreateRecipe(recipe model.Recipe) error {
	if recipe.Title == "" {
		return ErrTitleIsEmpty
	}
	return d.recipeRepo.CreateRecipe(recipe)
}

func (d *DefaultRecipeService) DeleteRecipe(id int) error {
	if id <= 0 {
		return ErrIDIsNotValid
	}

	err := d.recipeRepo.DeleteRecipe(id)
	if err != nil {
		if errors.Is(err, repository.ErrRecipeNotFound) {
			return ErrRecipeNotFound
		}
		return err
	}

	return nil
}