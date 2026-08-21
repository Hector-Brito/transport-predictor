package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"transport-predictor.com/v2/internal/driver"
	transportlog "transport-predictor.com/v2/internal/transportLog"
	"transport-predictor.com/v2/internal/vehicle"
	"transport-predictor.com/v2/internal/weather"
)

type Handlers struct {
	Driver      *driver.Handler
	Vehicle     *vehicle.Handler
	Weather     *weather.Handler
	TranportLog *transportlog.Handler
}

func (s *Server) RegisterRoutes(h *Handlers) {
	s.engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	v2 := s.engine.Group("/api/v2")

	{
		v2.GET("/driver/:id", h.Driver.GetOne)
		v2.GET("/driver/", h.Driver.GetAll)
		v2.POST("/driver/", h.Driver.Create)
		v2.PATCH("/driver/:id", h.Driver.Update)
		v2.DELETE("/driver/:id", h.Driver.Delete)
	}

	{
		v2.GET("/vehicle/:id", h.Vehicle.GetOne)
		v2.GET("/vehicle/", h.Vehicle.GetAll)
		v2.POST("/vehicle/", h.Vehicle.Create)
		v2.PATCH("/vehicle/:id", h.Vehicle.Update)
		v2.DELETE("/vehicle/:id", h.Vehicle.Delete)
	}

	{
		v2.GET("/weather/:id", h.Weather.GetOne)
		v2.GET("/weather/", h.Weather.GetAll)
		v2.POST("/weather/", h.Weather.Create)
		v2.PATCH("/weather/:id", h.Weather.Update)
		v2.DELETE("/weather/:id", h.Weather.Delete)
	}

	{
		v2.GET("/transportlog/:id", h.TranportLog.GetOne)
		v2.GET("/transportlog/", h.TranportLog.GetAll)
		v2.POST("/transportlog/", h.TranportLog.Create)
		v2.PATCH("/transportlog/:id", h.TranportLog.Update)
		v2.DELETE("/transportlog/:id", h.TranportLog.Delete)
	}
}
