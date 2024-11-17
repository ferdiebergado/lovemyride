package spareparts

import (
	"log/slog"
	"net/http"

	"github.com/ferdiebergado/lovemyride/internal/pkg/http/request"
	"github.com/ferdiebergado/lovemyride/internal/pkg/http/response"
)

type Handler struct {
	service Service
}

func NewSparePartsHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateSparePart(w http.ResponseWriter, r *http.Request) {
	params, err := request.JSON[CreateParams](r)

	if err != nil {
		response.ServerError(w, "decode json", err)
		return
	}

	sparePart, err := h.service.Create(r.Context(), params)

	if err != nil {
		slog.Error("create spare parts", "Error:", err)
		http.Error(w, "create spare parts", http.StatusInternalServerError)
		return
	}

	res := &response.APIResponse{
		Success: true,
		Data:    sparePart,
	}

	err = response.JSON(w, http.StatusCreated, res)

	if err != nil {
		response.ServerError(w, "json marshal", err)
	}
}

func (h *Handler) GetSparePart(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	spareParts, err := h.service.Find(r.Context(), id)

	if err != nil {
		response.ServerError(w, "find spare part", err)
		return
	}

	res := &response.APIResponse{
		Success: true,
		Data:    spareParts,
	}

	err = response.JSON(w, http.StatusOK, res)

	if err != nil {
		response.ServerError(w, "json marshal", err)
	}
}

func (h *Handler) GetAllSpareParts(w http.ResponseWriter, r *http.Request) {
	spareParts, err := h.service.GetAll(r.Context())

	if err != nil {
		response.ServerError(w, err.Error(), err)
		return
	}

	res := &response.APIResponse{
		Success: true,
		Data:    spareParts,
	}

	err = response.JSON(w, http.StatusOK, res)

	if err != nil {
		response.ServerError(w, "json marshal", err)
	}
}

func (h *Handler) UpdateSparePart(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	params, err := request.JSON[UpdateParams](r)

	if err != nil {
		response.ServerError(w, "decode json", err)
		return
	}

	err = h.service.Update(r.Context(), id, params)

	if err != nil {
		response.ServerError(w, "update sparepart", err)
		return
	}

	res := &response.APIResponse{
		Success: true,
	}

	err = response.JSON(w, http.StatusOK, res)

	if err != nil {
		response.ServerError(w, "json marshal", err)
	}
}

func (h *Handler) DeleteSparePart(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.service.Delete(r.Context(), id)

	if err != nil {
		response.ServerError(w, "delete sparepart", err)
		return
	}

	res := &response.APIResponse{
		Success: true,
	}

	err = response.JSON(w, http.StatusOK, res)

	if err != nil {
		response.ServerError(w, "json marshal", err)
	}
}
