package driver

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"transport-predictor.com/v2/domain"
)


type Handler struct {
	service *Service
	validate *validator.Validate
}

func NewHandler(s *Service) *Handler{
	return &Handler{
		service: s,
		validate: validator.New(),
	}
}

func (h *Handler) GetOne(ctx *gin.Context) {
	IDParam := ctx.Param("id")

	ID, err := strconv.Atoi(IDParam)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"error":"Bad Request", "message":"driver 'id' is required."})
	}
	driver, err := h.service.GetOne(ctx.Request.Context(),ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Internal Server Error", "message":"The server was unable to complete your request", "internal-message":err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, driver)
}

func (h *Handler) GetAll(ctx *gin.Context) {
	drivers, err := h.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":"Internal Server Error", "message":"The server was unable to complete your request", "internal-message":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, drivers)
}

func (h *Handler) Create(ctx *gin.Context) {
	var driver domain.Driver;
	if err := ctx.ShouldBindBodyWithJSON(&driver); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"error":"Bad Request", "message":"Request body could not be read properly.", "internal-message":err.Error()})
		return
	}
	if err :=  h.validate.Struct(driver); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"error":"Bad Request", "message":"Request body could not be read properly.", "internal-message":err.Error()})
		return
	}
	
	result, err := h.service.Create(ctx.Request.Context(),&driver)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Internal Server Error", "message":"The server was unable to complete your request", "internal-message":err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}


func (h *Handler) Update(ctx *gin.Context) {
	IDParam := ctx.Param("id")

	ID, err := strconv.Atoi(IDParam)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"error":"Bad Request", "message":"driver 'id' is required."})
	}
	var updateDriver domain.UpdateDriverParams;

	if err := ctx.ShouldBindBodyWithJSON(&updateDriver); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"error":"Bad Request", "message":"Request body could not be read properly.", "internal-message":err.Error()})
		return
	}

	if err :=  h.validate.Struct(updateDriver); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"error":"Bad Request", "message":"Request body could not be read properly.", "internal-message":err.Error()})
		return
	}
	result, err := h.service.Update(ctx.Request.Context(),ID, &updateDriver)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Internal Server Error", "message":"The server was unable to complete your request", "internal-message":err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)

}


func (h *Handler) Delete(ctx *gin.Context) {
	IDParam := ctx.Param("id")

	ID, err := strconv.Atoi(IDParam)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"error":"Bad Request", "message":"driver 'id' is required."})
	}

	result, err := h.service.Delete(ctx.Request.Context(),ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Internal Server Error", "message":"The server was unable to complete your request", "internal-message":err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}