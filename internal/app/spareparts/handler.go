package spareparts

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ferdiebergado/lovemyride/internal/pkg/db"
	"github.com/ferdiebergado/lovemyride/internal/pkg/http/request"
	"github.com/ferdiebergado/lovemyride/internal/pkg/http/response"
	"github.com/ferdiebergado/lovemyride/internal/pkg/logging"
	"github.com/ferdiebergado/lovemyride/internal/web/html"
)

type Handler struct {
	service Service
	logger  *slog.Logger
}

const path = "/spareparts"

func NewSparePartsHandler(service Service) *Handler {
	return &Handler{
		service: service,
		logger:  logging.CreateLogger(),
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
		h.logger.Error("create spare parts", "Error:", err)
		http.Error(w, "create spare parts", http.StatusInternalServerError)
		return
	}

	res := &response.APIResponse[SparePart]{
		Success: true,
		Message: "Sparepart created.",
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

	res := &response.APIResponse[SparePart]{
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

	res := &response.APIResponse[[]SparePart]{
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

	res := &response.APIResponse[any]{
		Success: true,
		Message: "Sparepart updated.",
	}

	err = response.JSON(w, http.StatusOK, res)

	if err != nil {
		response.ServerError(w, "json marshal", err)
	}
}

func (h *Handler) DeleteSparePart(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.service.Delete(r.Context(), id, db.SoftDelete)

	if err != nil {
		response.ServerError(w, "delete sparepart", err)
		return
	}

	res := &response.APIResponse[any]{
		Success: true,
	}

	err = response.JSON(w, http.StatusOK, res)

	if err != nil {
		response.ServerError(w, "json marshal", err)
	}
}

func (h *Handler) ListSpareParts(w http.ResponseWriter, r *http.Request) {
	spareparts, err := h.service.GetAll(r.Context())

	if err != nil {
		response.ServerError(w, "get all spareparts", err)
		return
	}

	headers := []response.TableHeader{
		{
			Label: "ID",
			Field: "id",
		},
		{
			Label: "Description",
			Field: "description",
		},
		{
			Label: "Maintenance Interval",
			Field: "maintenance_interval",
		},
	}

	urlData := fmt.Sprintf(`data-url="%s"`, path)

	headersJSON, err := json.Marshal(headers)

	if err != nil {
		response.ServerError(w, "json marshal", err)
		return
	}

	headersData := fmt.Sprintf(`data-headers='%s'`, string(headersJSON))

	dataJSON, err := json.Marshal(spareparts)

	if err != nil {
		response.ServerError(w, "json marshal", err)
		return
	}

	data := fmt.Sprintf(`data-data='%s'`, string(dataJSON))

	tableData := response.TableData{
		URL:     urlData,
		Data:    data,
		Headers: headersData,
	}

	html.Render(w, tableData, "partials/datatable.html", "pages/spareparts/index.html")
}

func (h *Handler) ShowCreateForm(w http.ResponseWriter, _ *http.Request) {
	formData := request.FormData{
		Action: "/api" + path,
	}

	html.Render(w, formData, "partials/forms/spareparts.html", "pages/spareparts/create.html")
}

func (h *Handler) ViewSparePart(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	sparepart, err := h.service.Find(r.Context(), id)

	if err != nil {
		if errors.Is(err, db.ErrModelNotFound) {
			html.Render(w, nil, "404.html")
			return
		}

		response.ServerError(w, "find sparepart: ", err)
		return
	}

	formData := request.FormData{
		Values: sparepart,
	}

	html.Render(w, formData, "pages/spareparts/view.html")
}

func (h *Handler) EditSparePart(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	sparepart, err := h.service.Find(r.Context(), id)

	if err != nil {
		if errors.Is(err, db.ErrModelNotFound) {
			html.Render(w, nil, "404.html")
			return
		}

		response.ServerError(w, "find sparepart: ", err)
		return
	}

	formData := request.FormData{
		Action: "/api" + path + "/" + sparepart.ID,
		Method: "PATCH",
		Values: sparepart,
	}

	html.Render(w, formData, "partials/forms/spareparts.html", "pages/spareparts/edit.html")
}
