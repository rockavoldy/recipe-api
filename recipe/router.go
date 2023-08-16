package recipe

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/oklog/ulid/v2"
	"github.com/rockavoldy/recipe-api/common"
)

var (
	ErrMissingQuery = errors.New("missing search query")
)

func Router() *chi.Mux {
	r := chi.NewMux()

	r.Get("/", listRecipeHandler)
	r.Get("/search", searchRecipeHandler)
	r.Get("/{recipeId}", getRecipeHandler)
	r.Post("/", createRecipeHandler)
	r.Put("/{recipeId}", updateRecipeHandler)
	r.Delete("/{recipeId}", deleteRecipeHandler)

	return r
}

func listRecipeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp, err := List(ctx)
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, err)
	}

	status := http.StatusOK
	common.WriteResponse(w, status, common.Response{
		Message: http.StatusText(status),
		Data:    resp,
	})
}

func searchRecipeHandler(w http.ResponseWriter, r *http.Request) {
	// `/search` endpoint can have some filter in the Query params
	//  `q` will be used to search the recipe's name
	// `material` will be used to filter by material's name, separated with comma
	// `category` will take the ID of the category
	query := r.URL.Query()
	var recipeQuery string
	if querySearch := query.Get("q"); querySearch != "" {
		recipeQuery = strings.ToLower(querySearch)
	}

	var category ulid.ULID
	if queryCategory := query.Get("category"); queryCategory != "" {
		category, _ = ulid.Parse(queryCategory)
	}

	var materials []string
	if queryMaterial := query.Get("material"); queryMaterial != "" {
		materials = strings.Split(strings.ToLower(queryMaterial), ",")
	}

	ctx := r.Context()
	resp, err := Search(ctx, recipeQuery, category, materials)
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	status := http.StatusOK
	common.WriteResponse(w, status, common.Response{
		Message: http.StatusText(status),
		Data:    resp,
	})
}

func getRecipeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	recipeId := chi.URLParam(r, "recipeId")
	id, err := ulid.Parse(recipeId)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := Find(ctx, id)
	if err != nil {
		if err == ErrNotFound {
			common.WriteError(w, http.StatusNotFound, err)
			return
		}

		common.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	status := http.StatusOK
	common.WriteResponse(w, status, common.Response{
		Message: http.StatusText(status),
		Data:    resp,
	})
}

func createRecipeHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		common.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	ctx := r.Context()
	var data recipeJsonReq
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	id, err := Create(ctx, data)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}

	status := http.StatusOK
	common.WriteResponse(w, status, common.Response{
		Message: http.StatusText(status),
		Data:    id,
	})
}

func updateRecipeHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		common.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	recipeId := chi.URLParam(r, "recipeId")
	id, err := ulid.Parse(recipeId)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}

	ctx := r.Context()
	var data recipeJsonReq
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	resp, err := Update(ctx, id, data)
	if err != nil {
		if err == ErrNotFound {
			common.WriteError(w, http.StatusNotFound, err)
			return
		}

		common.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	status := http.StatusOK
	common.WriteResponse(w, status, common.Response{
		Message: http.StatusText(status),
		Data:    resp,
	})
}

func deleteRecipeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	recipeId := chi.URLParam(r, "recipeId")
	id, err := ulid.Parse(recipeId)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = Delete(ctx, id)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
	}

	status := http.StatusOK
	common.WriteResponse(w, status, common.Response{
		Message: http.StatusText(status),
		Data:    nil,
	})
}
