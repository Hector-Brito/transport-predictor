package transportlog

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"transport-predictor.com/v2/domain"
)

type Handler struct {
	service  *Service
	validate *validator.Validate
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *Handler) GetOne(ctx *gin.Context) {
	IDParam := ctx.Param("id")

	ID, err := strconv.Atoi(IDParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "transport_log 'id' is required."})
		return
	}

	transport_log, err := h.service.GetOne(ctx.Request.Context(), ID)

	if err != nil {

		if errors.Is(err, domain.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Not found", "message": domain.ErrNotFound.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "The server was unable to complete your request", "internal-message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transport_log)
}

func (h *Handler) GetAll(ctx *gin.Context) {
	transport_logs, err := h.service.GetAll(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "The server was unable to complete your request", "internal-message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transport_logs)
}

func (h *Handler) Create(ctx *gin.Context) {
	var transport_log domain.TransportLog

	if err := ctx.ShouldBindBodyWithJSON(&transport_log); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Request body could not be read properly.", "internal-message": err.Error()})
		return
	}

	if err := h.validate.Struct(transport_log); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Request body could not be read properly.", "internal-message": err.Error()})
		return
	}

	result, err := h.service.Create(ctx.Request.Context(), &transport_log)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "The server was unable to complete your request", "internal-message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (h *Handler) Update(ctx *gin.Context) {
	IDParam := ctx.Param("id")

	ID, err := strconv.Atoi(IDParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "vehicle 'id' is required."})
		return
	}

	var updateTransportLogParams domain.UpdateTransportLogParams

	if err := ctx.ShouldBindBodyWithJSON(&updateTransportLogParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Request body could not be read properly.", "internal-message": err.Error()})
		return
	}

	if err := h.validate.Struct(updateTransportLogParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Request body could not be read properly.", "internal-message": err.Error()})
		return
	}

	if updateTransportLogParams.DepartureDate == nil && updateTransportLogParams.ArrivalDate == nil && updateTransportLogParams.Latency == nil && updateTransportLogParams.DayOfWeek == nil && updateTransportLogParams.IsFortnight == nil && updateTransportLogParams.Observations == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Required request body is missing."})
		return
	}

	result, err := h.service.Update(ctx.Request.Context(), ID, &updateTransportLogParams)

	if err != nil {

		if errors.Is(err, domain.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Not found", "message": domain.ErrNotFound.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "The server was unable to complete your request", "internal-message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) Delete(ctx *gin.Context) {

	IDParam := ctx.Param("id")

	ID, err := strconv.Atoi(IDParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "transport_log 'id' is required."})
		return
	}

	result, err := h.service.Delete(ctx.Request.Context(), ID)

	if err != nil {

		if errors.Is(err, domain.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Not found", "message": domain.ErrNotFound.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "The server was unable to complete your request", "internal-message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
