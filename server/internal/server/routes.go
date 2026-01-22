package server

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"transport-predictor.com/v2/internal/driver"
)

type Handlers struct {
	Driver *driver.Handler
}

func (s *Server) RegisterRoutes(h *Handlers) {
	s.engine.GET("/",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":"OK",
		})
	})

	v2 := s.engine.Group("/api/v2")

	{
		v2.GET("/driver/:id", h.Driver.GetOne)
		v2.GET("/driver", h.Driver.GetAll)
		v2.POST("/driver", h.Driver.Create)
		v2.PATCH("/driver/:id", h.Driver.Update)
		v2.DELETE("/driver/:id",h.Driver.Delete)
	}
}