package vehicle

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
	//Obtener parametros de url de la peticion
	IDParam := ctx.Param("id")

	ID, err := strconv.Atoi(IDParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "vehicle 'id' is required."})
		return
	}

	vehicle, err := h.service.GetOne(ctx.Request.Context(), ID)

	if err != nil {

		if errors.Is(err, domain.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Not found", "message": domain.ErrNotFound.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "The server was unable to complete your request", "internal-message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, vehicle)
}

func (h *Handler) GetAll(ctx *gin.Context) {
	vehicles, err := h.service.GetAll(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "The server was unable to complete your request", "internal-message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, vehicles)
}

func (h *Handler) Create(ctx *gin.Context) {
	var vehicle domain.Vehicle

	//Enlaza los datos de la peticion con la variable "vehicle".
	if err := ctx.ShouldBindBodyWithJSON(&vehicle); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Request body could not be read properly.", "internal-message": err.Error()})
		return
	}

	if err := h.validate.Struct(vehicle); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Request body could not be read properly.", "internal-message": err.Error()})
		return
	}

	result, err := h.service.Create(ctx.Request.Context(), &vehicle)

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

	var updateVehicle domain.UpdateVehicleParams

	if err := ctx.ShouldBindBodyWithJSON(&updateVehicle); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Request body could not be read properly.", "internal-message": err.Error()})
		return
	}

	if err := h.validate.Struct(updateVehicle); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Request body could not be read properly.", "internal-message": err.Error()})
		return
	}

	if updateVehicle.Name == nil && updateVehicle.NickName == nil && updateVehicle.LicensePlate == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Required request body is missing."})
		return
	}
	result, err := h.service.Update(ctx.Request.Context(), ID, &updateVehicle)

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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "vehicle 'id' is required."})
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
