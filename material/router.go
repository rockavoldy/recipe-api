package material

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oklog/ulid/v2"
	"github.com/rockavoldy/recipe-api/common"
)

func Router() *chi.Mux {
	r := chi.NewMux()

	r.Get("/", listMaterialHandler)
	r.Get("/{materialId}", getMaterialHandler)
	r.Post("/", createMaterialHandler)
	r.Put("/{materialId}", updateMaterialHandler)
	r.Delete("/{materialId}", deleteMaterialHandler)

	return r
}

func listMaterialHandler(w http.ResponseWriter, r *http.Request) {
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

func getMaterialHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	materialId := chi.URLParam(r, "materialId")
	id, err := ulid.Parse(materialId)
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

func createMaterialHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		common.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	ctx := r.Context()
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	r.Body.Close()
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}
	id, err := Create(ctx, data["name"])
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

func updateMaterialHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		common.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	ctx := r.Context()
	materialId := chi.URLParam(r, "materialId")
	id, err := ulid.Parse(materialId)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}

	var data map[string]string
	err = json.NewDecoder(r.Body).Decode(&data)
	r.Body.Close()
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}
	resp, err := Update(ctx, id, data["name"])
	if err != nil {
		if err == ErrNotFound {
			common.WriteError(w, http.StatusNotFound, err)
			return
		}

		common.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}

	status := http.StatusOK
	common.WriteResponse(w, status, common.Response{
		Message: http.StatusText(status),
		Data:    resp,
	})
}

func deleteMaterialHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	materialId := chi.URLParam(r, "materialId")
	id, err := ulid.Parse(materialId)
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
