package handler

import (
	"encoding/json"
	"net/http"
	"RecipeCollection/service"
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