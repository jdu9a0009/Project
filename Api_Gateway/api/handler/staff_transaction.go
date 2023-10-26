package handler

import (
	"net/http"
	"strconv"

	sale_service "api-gateway-service/genproto/sale_service"

	"github.com/gin-gonic/gin"
)

// CreateStaffTransaction godoc
// @Router       /v1/staff-transactions [post]
// @Summary      Create a new staff-transaction
// @Description  Create a new staff-transaction with the provided details
// @Tags         staff-transactions
// @Accept       json
// @Produce      json
// @Param        staff_transaction     body  sale_service.CreateStaffTransactionRequest true  "data of the staff transaction"
// @Success      201  {object}  sale_service.IdResponse
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) CreateStaffTransaction(ctx *gin.Context) {
	var staffTr = sale_service.StaffTransaction{}

	if err := ctx.ShouldBindJSON(&staffTr); err != nil {
		h.handlerResponse(ctx, "CreateStaffTransaction", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.StaffTransactionService().Create(ctx, &sale_service.CreateStaffTransactionRequest{
		Id:         staffTr.Id,
		SaleId:     staffTr.SaleId,
		StaffId:    staffTr.StaffId,
		Type:       staffTr.Type,
		SourceType: staffTr.SourceType,
		Amount:     staffTr.Amount,
		AboutText:  staffTr.AboutText,
	})

	if err != nil {
		h.handlerResponse(ctx, "StaffTransactionService().Create", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "create staff transaction response", http.StatusOK, resp)
}

// ListStaffTransaction godoc
// @Router       /v1/staff-transactions [get]
// @Summary      List staff-transactions
// @Description  get staff-transactions
// @Tags         staff-transactions
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"  Default(10)
// @Param		 page     query     int  false  "page for response"   Default(1)
// @Param        amount_from     query     string false "search by amount_from"
// @Param        amount_to     query     string false "search by amount_to"
// @Param        sale_id     query     string false "search by sale_id"
// @Param        type     query     string false "search by type"
// @Param        staff_id     query     string false "search by staff_id"
// @Success      200  {array}   sale_service.StaffTransaction
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetAllStaffTransaction(ctx *gin.Context) {
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

	amount, err := strconv.ParseFloat(ctx.Query("amount"), 64)
	if err != nil {
		h.handlerResponse(ctx, "error get amount_from", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.StaffTransactionService().GetAll(ctx.Request.Context(), &sale_service.GetAllStaffTransactionRequest{
		Offset:  int32(page),
		Limit:   int32(limit),
		Type:    ctx.Query("type"),
		SaleId:  ctx.Query("sale_id"),
		StaffId: ctx.Query("staff_id"),
		Amount:  float32(amount),
	})

	if err != nil {
		h.handlerResponse(ctx, "error GetAllStaffTransaction", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get list staff transaction response", http.StatusOK, resp)
}

// GetStaffTransaction godoc
// @Router       /v1/staff-transactions/{id} [get]
// @Summary      Get a staff-transaction by ID
// @Description  Retrieve a staff-transaction by its unique identifier
// @Tags         staff-transactions
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "transaction ID to retrieve"
// @Success      200  {object}  sale_service.StaffTransaction
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetStaffTransaction(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.StaffTransactionService().Get(ctx.Request.Context(), &sale_service.IdRequest{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error staff transaction GetById", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get staff transaction response", http.StatusOK, resp)
}

// UpdateStaffTransaction godoc
// @Router       /v1/staff-transactions/{id} [put]
// @Summary      Update an existing staff-transaction
// @Description  Update an existing staff-transaction with the provided details
// @Tags         staff-transactions
// @Accept       json
// @Produce      json
// @Param        id       path    string     true    "Transaction ID to update"
// @Param        staff_transaction   body    sale_service.UpdateStaffTransactionRequest  true    "Updated data for the sale"
// @Success      200  {object}  sale_service.UpdateStaffTransactionRequest
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) UpdateStaffTransaction(ctx *gin.Context) {
	var staffTr = sale_service.StaffTransaction{}
	staffTr.Id = ctx.Param("id")
	err := ctx.ShouldBindJSON(&staffTr)
	if err != nil {
		h.handlerResponse(ctx, "error while binding", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.StaffTransactionService().Update(ctx.Request.Context(), &sale_service.UpdateStaffTransactionRequest{
		Id:         staffTr.Id,
		SaleId:     staffTr.SaleId,
		StaffId:    staffTr.StaffId,
		Type:       staffTr.Type,
		SourceType: staffTr.SourceType,
		Amount:     staffTr.Amount,
		AboutText:  staffTr.AboutText,
	})

	if err != nil {
		h.handlerResponse(ctx, "error staff transaction Update", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "update staff transaction response", http.StatusOK, resp)
}

// DeleteStaffTransaction godoc
// @Router       /v1/staff-transactions/{id} [delete]
// @Summary      Delete a staff-transaction
// @Description  delete a staff-transaction by its unique identifier
// @Tags         staff-transactions
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Transaction ID to retrieve"
// @Success      200  {object}  sale_service.IdRequest
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) DeleteStaffTransaction(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.StaffTransactionService().Delete(ctx.Request.Context(), &sale_service.IdRequest{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error staff transaction Delete", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "delete staff transaction response", http.StatusOK, resp)
}
