package unit

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oklog/ulid/v2"
	"github.com/rockavoldy/recipe-api/common"
)

func Router() *chi.Mux {
	r := chi.NewMux()

	r.Get("/", listUnitHandler)
	r.Get("/{unitId}", getUnitHandler)
	r.Post("/", createUnitHandler)
	r.Put("/{unitId}", updateUnitHandler)
	r.Delete("/{unitId}", deleteUnitHandler)

	return r
}

func listUnitHandler(w http.ResponseWriter, r *http.Request) {
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

func getUnitHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	unitId := chi.URLParam(r, "unitId")
	id, err := ulid.Parse(unitId)
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

func createUnitHandler(w http.ResponseWriter, r *http.Request) {
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

func updateUnitHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		common.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	ctx := r.Context()
	unitId := chi.URLParam(r, "unitId")
	id, err := ulid.Parse(unitId)
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

func deleteUnitHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	unitId := chi.URLParam(r, "unitId")
	id, err := ulid.Parse(unitId)
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
