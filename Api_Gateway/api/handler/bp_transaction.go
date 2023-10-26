package handler

import (
	"net/http"
	"strconv"

	sale_service "api-gateway-service/genproto/sale_service"

	"github.com/gin-gonic/gin"
)

// CreateBranchProductTransaction godoc
// @Router       /v1/branch-product-transactions [post]
// @Summary      Create a new branch-product-transaction
// @Description  Create a new branch-product-transaction with the provided details
// @Tags         branch-product-transactions
// @Accept       json
// @Produce      json
// @Param        br_pr_tr     body  sale_service.BrPrTrCreateReq  true  "data of the staff transaction"
// @Success      201  {object}  sale_service.BrPrTrCreateResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) CreateBrPrTr(ctx *gin.Context) {
	var brPrTr = sale_service.BpTransaction{}

	err := ctx.ShouldBindJSON(&brPrTr)
	if err != nil {
		h.handlerResponse(ctx, "CreateBrPrTr", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.BranchProductTransactionService().Create(ctx, &sale_service.CreateBpTransactionRequest{
		Id:        brPrTr.Id,
		BranchId:  brPrTr.BranchId,
		StaffId:   brPrTr.StaffId,
		ProductId: brPrTr.ProductId,
		Price:     brPrTr.Price,
		Type:      brPrTr.Type,
		Quantity:  brPrTr.Quantity,
	})

	if err != nil {
		h.handlerResponse(ctx, "BranchProductTransactionService().Create", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "create branch_product_transaction response", http.StatusOK, resp)
}

// ListBranchProductTransaction godoc
// @Router       /v1/branch-product-transactions [get]
// @Summary      List branch-product-transactions
// @Description  get branch-product-transactions
// @Tags         branch-product-transactions
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"  Default(10)
// @Param		 page     query     int  false  "page for response"   Default(1)
// @Param        branch_id     query     string false "search by branch_id"
// @Param        type     query     string false "search by type"
// @Param        staff_id     query     string false "search by staff_id"
// @Success      200  {array}   sale_service.BrPrTransaction
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetAllBrPrTr(ctx *gin.Context) {
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

	resp, err := h.services.BranchProductTransactionService().GetAll(ctx.Request.Context(), &sale_service.GetAllBpTransactionRequest{
		Offset: int32(page),
		Limit:  int32(limit),
		Search: ctx.Query("type"),
	})

	if err != nil {
		h.handlerResponse(ctx, "error GetAllBrPrTr", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get All branch_product_transaction response", http.StatusOK, resp)
}

// GetBranchProductTransaction godoc
// @Router       /v1/branch-product-transactions/{id} [get]
// @Summary      Get a branch-product-transaction by ID
// @Description  Retrieve a branch-product-transaction by its unique identifier
// @Tags         branch-product-transactions
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "br_pr_transaction ID to retrieve"
// @Success      200  {object}  sale_service.BrPrTransaction
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetBrPrTr(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.BranchProductTransactionService().Get(ctx.Request.Context(), &sale_service.IdRequest{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error branchProductTransaction GetById", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get branch_product_transaction response", http.StatusOK, resp)
}

// UpdateBranchProductTransaction godoc
// @Router       /v1/branch-product-transactions/{id} [put]
// @Summary      Update an existing branch-product-transaction
// @Description  Update an existing branch-product-transaction with the provided details
// @Tags         branch-product-transactions
// @Accept       json
// @Produce      json
// @Param        id       path    string     true    "br_pr_transaction ID to update"
// @Param        br_pr_transaction   body    sale_service.BrPrTrUpdateReq  true    "Updated data for the br_pr_tr"
// @Success      200  {object}  sale_service.BrPrTrUpdateResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) UpdateBrPrTr(ctx *gin.Context) {
	var brPrTr = sale_service.BpTransaction{}
	brPrTr.Id = ctx.Param("id")
	err := ctx.ShouldBindJSON(&brPrTr)
	if err != nil {
		h.handlerResponse(ctx, "error while binding", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.BranchProductTransactionService().Update(ctx.Request.Context(), &sale_service.UpdateBpTransactionRequest{
		Id:        brPrTr.Id,
		BranchId:  brPrTr.BranchId,
		StaffId:   brPrTr.StaffId,
		ProductId: brPrTr.ProductId,
		Price:     brPrTr.Price,
		Type:      brPrTr.Type,
		Quantity:  brPrTr.Quantity,
	})

	if err != nil {
		h.handlerResponse(ctx, "error branch_product_transaction Update", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "update branch_product_transaction response", http.StatusOK, resp)
}

// DeleteBranchProductTransaction godoc
// @Router       /v1/branch-product-transactions/{id} [delete]
// @Summary      Delete a branch-product-transaction
// @Description  delete a branch-product-transaction by its unique identifier
// @Tags         branch-product-transactions
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "br_pr_transaction ID to retrieve"
// @Success      200  {object}  sale_service.BrPrTrDeleteResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) DeleteBrPrTr(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.BranchProductTransactionService().Delete(ctx.Request.Context(), &sale_service.IdRequest{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error branch_product_transaction Delete", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "delete branch_product_transaction response", http.StatusOK, resp)
}
