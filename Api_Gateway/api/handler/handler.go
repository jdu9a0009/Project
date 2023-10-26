package handler

import (
	"api-gateway-service/config"
	"api-gateway-service/pkg/logger"
	"api-gateway-service/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg      config.Config
	log      logger.LoggerI
	services services.ServiceManagerI
}

func NewHandler(cfg config.Config, log logger.LoggerI, srvc services.ServiceManagerI) *Handler {
	return &Handler{
		cfg:      cfg,
		log:      log,
		services: srvc,
	}
}

type Response struct {
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func (h *Handler) handlerResponse(c *gin.Context, path string, code int, message interface{}) {
	response := Response{
		Status: code,
		Data:   message,
	}

	switch {
	case code < 300:
		h.log.Info(path, logger.Any("info", response))
	case code >= 400:
		h.log.Error(path, logger.Any("error", response))
	}

	c.JSON(code, response)
}
