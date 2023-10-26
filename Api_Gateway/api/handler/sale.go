package handler

import (
	"net/http"
	"strconv"

	sale_service "api-gateway-service/genproto/sale_service"

	"github.com/gin-gonic/gin"
)

// CreateSale godoc
// @Router       /v1/sales [post]
// @Summary      Create a new sale
// @Description  Create a new sale with the provided details
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        sale     body  sale_service.CreateSaleRequest  true  "data of the sale"
// @Success      201  {object}  sale_service.IdResponse
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) CreateSale(ctx *gin.Context) {
	var sale = sale_service.Sale{}

	err := ctx.ShouldBindJSON(&sale)
	if err != nil {
		h.handlerResponse(ctx, "CreateSale", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.SaleService().Create(ctx, &sale_service.CreateSaleRequest{
		Id:               sale.Branch_Id,
		Branch_Id:        sale.Branch_Id,
		ShopAssistant_Id: sale.ShopAssistant_Id,
		CashierId:        sale.CashierId,
		Price:            sale.Price,
		PaymentType:      sale.PaymentType,
		ClientName:       sale.ClientName,
	})

	if err != nil {
		h.handlerResponse(ctx, "SaleService().Create", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "create sale response", http.StatusOK, resp)
}

// GetAllSales godoc
// @Router       /v1/sales [get]
// @Summary      GetAll sales
// @Description  get sales
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"  Default(10)
// @Param		 page     query     int  false  "page for response"   Default(1)
// @Param        branch_id     query     string false "search by branch_id"
// @Param        client_name     query     string false "search by client_name"
// @Param        payment_type     query     string false "search by payment_type"
// @Param        shop_assistant_id     query     string false "search by shop_assistant_id"
// @Param        cashier_id     query     string false "search by cashier_id"
// @Param        status     query     string false "search by status"
// @Param        created_at_from     query     string false "search by created_at_from"
// @Param        created_at_to     query     string false "search by created_at_to"
// @Success      200  {array}   sale_service.Sale
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetAllSale(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		h.handlerResponse(ctx, "error get page", http.StatusBadRequest, err.Error())
		return
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		h.handlerResponse(ctx, "error get limit", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.SaleService().GetAll(ctx.Request.Context(), &sale_service.GetAllSaleRequest{
		Offset:           int32(page),
		Limit:            int32(limit),
		Search:           ctx.Query("search"),
		Branch_Id:        ctx.Query("branch_id"),
		PaymentType:      ctx.Query("payment_type"),
		ShopAssistant_Id: ctx.Query("shop_assistant_id"),
		CashierId:        ctx.Query("cashier_id"),
		CreatedAtFrom:    ctx.DefaultQuery("created_at_from", "2000-01-01"),
		CreatedAtTo:      ctx.DefaultQuery("created_at_to", "2095-12-12"),
	})

	if err != nil {
		h.handlerResponse(ctx, "error GetListSale", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get list sale response", http.StatusOK, resp)
}

// GetSale godoc
// @Router       /v1/sales/{id} [get]
// @Summary      Get a sale by ID
// @Description  Retrieve a sale by its unique identifier
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Sale ID to retrieve"
// @Success      200  {object}  sale_service.Sale
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetSale(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.SaleService().Get(ctx.Request.Context(), &sale_service.IdRequest{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error sale GetById", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get sale response", http.StatusOK, resp)
}

// UpdateSale godoc
// @Router       /v1/sales/{id} [put]
// @Summary      Update an existing sale
// @Description  Update an existing sale with the provided details
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        id       path    string     true    "Sale ID to update"
// @Param        sale   body    sale_service.UpdateSaleRequest true    "Updated data for the sale"
// @Success      200  {object}  Response{data=string}
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) UpdateSale(ctx *gin.Context) {
	var sale = sale_service.Sale{}
	sale.Id = ctx.Param("id")
	err := ctx.ShouldBindJSON(&sale)
	if err != nil {
		h.handlerResponse(ctx, "error while binding", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.SaleService().Update(ctx.Request.Context(), &sale_service.UpdateSaleRequest{
		Id:               sale.Id,
		Branch_Id:        sale.Branch_Id,
		ShopAssistant_Id: sale.ShopAssistant_Id,
		CashierId:        sale.CashierId,
		Price:            sale.Price,
		PaymentType:      sale.PaymentType,
		Status:           sale.Status,
		ClientName:       sale.ClientName,
	})

	if err != nil {
		h.handlerResponse(ctx, "error sale Update", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "update sale response", http.StatusOK, resp)
}

// DeleteSale godoc
// @Router       /v1/sales/{id} [delete]
// @Summary      Delete a sale
// @Description  delete a sale by its unique identifier
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Sale ID to retrieve"
// @Success      200  {object}  sale_service.IdResponse
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) DeleteSale(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.SaleService().Delete(ctx.Request.Context(), &sale_service.IdRequest{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error sale Delete", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "delete sale response", http.StatusOK, resp)
}
