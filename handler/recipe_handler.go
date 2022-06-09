package handler

import (
	"RecipeCollection/service"
	"RecipeCollection/model"
	"encoding/json"
	"net/http"
	"strconv"
	"errors"

	"github.com/julienschmidt/httprouter"
)

type recipeHandler struct {
	service service.IRecipeService
}

func NewRecipeHandler(rs service.IRecipeService) *recipeHandler {
	return &recipeHandler{service: rs}
}

func (rh *recipeHandler) GetRecipes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	recipes, err := rh.service.GetRecipes()
	if err != nil {
		http.Error(w, "Unable to get all recipes", http.StatusInternalServerError)
		return
	}

	jsonStr, err := json.Marshal(recipes)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}

func(rh *recipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))

	recipe, err := rh.service.GetRecipe(id)
	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if errors.Is(err, service.ErrRecipeNotFound) { // Test yaz
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonStr, err := json.Marshal(recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}

func (rh *recipeHandler) CreateRecipe(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var recipe model.Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, "error when decoding json", http.StatusInternalServerError)
		return
	}

	err = rh.service.CreateRecipe(recipe)
	if err != nil {
		if errors.Is(err, service.ErrTitleIsEmpty) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Recipe is successfully created"))
}

func (rh *recipeHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))

	err := rh.service.DeleteRecipe(id)
	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) || errors.Is(err, service.ErrRecipeNotFound) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Recipe has been deleted succesfully"))
}