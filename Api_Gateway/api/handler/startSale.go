package handler

import (
	sale_service "api-gateway-service/genproto/sale_service"
	"api-gateway-service/pkg/logger"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateSale godoc
// @Router       /start [POST]
// @Summary      Create Sale
// @Description  Create Sale
// @Tags         START
// @Accept       json
// @Produce      json
// @Param        data  body      models.StartSale  true  "Sale data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) StartSale(c *gin.Context) {

	var sale models.StartSale
	err := c.ShouldBindJSON(&sale)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.services.SaleService().Create(c.Request.Context(), &sale_service.CreateSaleRequest{
		BranchId:        sale.BranchId,
		ShopAssistentId: sale.ShopAssistentId,
		CashierId:       sale.CashierId,
		PaymentType:     sale.PaymentType,
		Price:           0.0,
		Status:          "start",
		ClientName:      sale.ClientName,
	})
	if err != nil {
		fmt.Println("error Sale Create:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusCreated, response.CreateResponse{Id: resp.GetId()})

}
