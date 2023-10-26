package handler

import (
	"net/http"
	"strconv"

	branch_service "api-gateway-service/genproto/branch_service"

	"github.com/gin-gonic/gin"
)

// CreateBranchProduct godoc
// @Router       /v1/branch-products [post]
// @Summary      Create a new branch_product
// @Description  Create a new branch_product with the provided details
// @Tags         branch-products
// @Accept       json
// @Produce      json
// @Param        branch_product     body  branch_service.BranchProductCreateReq  true  "data of the branch_product"
// @Success      201  {object}  branch_service.BranchProductCreateResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) CreateBranchProduct(ctx *gin.Context) {
	var branchPr = branch_service.BranchProduct{}

	err := ctx.ShouldBindJSON(&branchPr)
	if err != nil {
		h.handlerResponse(ctx, "CreateBranchProduct", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.BranchProductService().Create(ctx, &branch_service.CreateBranchProductRequest{
		ProductId: branchPr.ProductId,
		BranchId:  branchPr.BranchId,
		Count:     branchPr.Count,
	})

	if err != nil {
		h.handlerResponse(ctx, "BranchProductService().Create", http.StatusBadRequest, err.Error())

		return
	}

	h.handlerResponse(ctx, "create branch product response", http.StatusOK, resp)
}

// ListBranchProducts godoc
// @Router       /v1/branch-products [get]
// @Summary      List branch_products
// @Description  get branch_products
// @Tags         branch-products
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"  Default(10)
// @Param		 page     query     int  false  "page for response"   Default(1)
// @Param        branch_id     query     string false "search by branch_id"
// @Success      200  {array}   branch_service.BranchProduct
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetListBranchProduct(ctx *gin.Context) {
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

	resp, err := h.services.BranchProductService().GetAll(ctx.Request.Context(), &branch_service.GetAllBranchProductRequest{
		Offset: int32(page),
		Limit:  int32(limit),
		Search: ctx.Query("branch_id"),
	})

	if err != nil {
		h.handlerResponse(ctx, "error GetAllBranchProduct", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get list branch product response", http.StatusOK, resp)
}

// GetBranchProduct godoc
// @Router       /v1/branch-products/{id} [get]
// @Summary      Get a branch_product by product_id
// @Description  Retrieve a branch_product by its product_id
// @Tags         branch-products
// @Accept       json
// @Produce      json
// @Param        product_id   path    string     true    "product_id to retrieve"
// @Success      200  {object}  branch_service.BranchProduct
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetBranchProduct(ctx *gin.Context) {
	id := ctx.Param("product_id")

	resp, err := h.services.BranchProductService().Get(ctx.Request.Context(), &branch_service.IdRequest{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error branch product GetById", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get branch product response", http.StatusOK, resp)
}

// UpdateBranchProduct godoc
// @Router       /v1/branch-products/{id} [put]
// @Summary      Update an existing branch_product
// @Description  Update an existing branch_product with the provided details
// @Tags         branch-products
// @Accept       json
// @Produce      json
// @Param        product_id       path    string     true    "product_id to update"
// @Param        branch_product   body    branch_service.BranchProductUpdateReq  true    "Updated data for the branch_product"
// @Success      200  {object}  branch_service.BranchProductUpdateResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) UpdateBranchProduct(ctx *gin.Context) {
	var branchPr = branch_service.BranchProduct{}
	branchPr.ProductId = ctx.Param("product_id")
	err := ctx.ShouldBindJSON(&branchPr)
	if err != nil {
		h.handlerResponse(ctx, "error while binding", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.BranchProductService().Update(ctx.Request.Context(), &branch_service.UpdateBranchProductRequest{
		ProductId: branchPr.ProductId,
		BranchId:  branchPr.BranchId,
		Count:     int32(branchPr.Count),
	})

	if err != nil {
		h.handlerResponse(ctx, "error branch Update", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "update branch response", http.StatusOK, resp)
}

// DeleteBranchProduct godoc
// @Router       /v1/branch-products/{id} [delete]
// @Summary      Delete a branch_product
// @Description  delete a branch_product by its product_id
// @Tags         branch-products
// @Accept       json
// @Produce      json
// @Param        product_id   path    string     true    "product_id to retrieve"
// @Success      200  {object}  branch_service.BranchProductDeleteResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) DeleteBranchProduct(ctx *gin.Context) {
	id := ctx.Param("product_id")

	resp, err := h.services.BranchProductService().Delete(ctx.Request.Context(), &branch_service.IdRequest{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error branch product Delete", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "delete branch product response", http.StatusOK, resp)
}
