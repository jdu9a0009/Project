package handler

import (
	"fmt"
	"net/http"
	"strconv"

	sale_service "api-gateway-service/genproto/sale_service"

	"github.com/gin-gonic/gin"
)

// CreateSaleProduct godoc
// @Router       /v1/sale-products [post]
// @Summary      Create a new sale-product
// @Description  Create a new sale-product with the provided details
// @Tags         sale-products
// @Accept       json
// @Produce      json
// @Param        sale-product     body  sale_service.CreateSaleProductRequest  true  "data of the sale-product"
// @Success      201  {object}  sale_service.IdResponse
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) CreateSaleProduct(ctx *gin.Context) {
	var saleProduct = sale_service.SaleProduct{}

	err := ctx.ShouldBindJSON(&saleProduct)
	if err != nil {
		h.handlerResponse(ctx, "CreateSaleProduct", http.StatusBadRequest, err.Error())
		return
	}

	respSale, err := h.services.SaleService().Get(ctx, &sale_service.IdRequest{Id: saleProduct.SaleId})
	if err != nil {
		h.handlerResponse(ctx, "SaleGet error", http.StatusInternalServerError, err.Error())
		return
	}

	respWarehouse, err := h.services.BranchProductService().GetProduct(ctx, &sale_service.GetProductRequest{
		BranchId:  respSale.BranchId,
		ProductId: saleProduct.ProductId,
	})

	if err != nil {
		h.handlerResponse(ctx, "Check product in warehouse", http.StatusInternalServerError, err.Error())
		return
	}
	respSaleProduct, err := h.services.SaleProductService().GetSaleById(ctx, &sale_service.SaleIdRequest{
		SaleId:    saleProduct.SaleId,
		ProductId: saleProduct.ProductId,
	})
	if err != nil {
		if respWarehouse.Count >= int32(saleProduct.Quantity) {
			resp, err := h.services.SaleProductService().Create(ctx, &sale_service.CreateSaleProductRequest{
				Id:        saleProduct.Id,
				SaleId:    saleProduct.SaleId,
				ProductId: saleProduct.ProductId,
				Quantity:  saleProduct.Quantity,
				Price:     saleProduct.Price,
			})
			if err != nil {
				h.handlerResponse(ctx, "SaleProductService().Create", http.StatusBadRequest, err.Error())

				return
			}

			h.handlerResponse(ctx, "create sale product response", http.StatusOK, resp)
		} else {
			fmt.Println("Not Enough Product", err.Error())
			ctx.JSON(http.StatusBadRequest, "not enough product")
			return
		}

	}
	if respWarehouse.Count >= int32(saleProduct.Quantity)+respWarehouse.Quantity {
		resp, err := h.services.SaleProductService().Create(ctx, &sale_service.CreateSaleProductRequest{
			Id:        saleProduct.Id,
			SaleId:    saleProduct.SaleId,
			ProductId: saleProduct.ProductId,
			Quantity:  saleProduct.Quantity + int32(saleProduct.Quantity),
			Price:     saleProduct.Price,
		})
		if err != nil {
			h.handlerResponse(ctx, "SaleProductService().Update", http.StatusBadRequest, err.Error())

			return
		}
		ctx.JSON(http.StatusOK, resp)

	} else {
		fmt.Println("Not Enough Produtxt", err.Error())
		ctx.JSON(http.StatusBadRequest, "not enough product")
		return
	}

}

// GetAllSaleProducts godoc
// @Router       /v1/sale-products [get]
// @Summary      GetAll sale-products
// @Description  get sale-products
// @Tags         sale-products
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"  Default(10)
// @Param		 page     query     int  false  "page for response"   Default(1)
// @Success      200  {array}   sale_service.SaleProduct
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetAllSaleProduct(ctx *gin.Context) {
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

	resp, err := h.services.SaleProductService().GetAll(ctx.Request.Context(), &sale_service.GetAllSaleProductRequest{
		Page:  int32(page),
		Limit: int32(limit),
	})

	if err != nil {
		h.handlerResponse(ctx, "error GetAllSaleProduct", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get All sale product response", http.StatusOK, resp)
}

// GetSaleProduct godoc
// @Router       /v1/sale-products/{id} [get]
// @Summary      Get a sale_product by ID
// @Description  Retrieve a sale by its unique identifier
// @Tags         sale-products
// @Accept       json
// @Produce      json
// @Param        sale_id   path    string     true    "SaleProduct ID to retrieve"
// @Success      200  {object}  sale_service.SaleProduct
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetSaleProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.SaleProductService().Get(ctx.Request.Context(), &sale_service.IdRequest{
		Id: id,
	})
	if err != nil {
		h.handlerResponse(ctx, "error sale GetById", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get sale product response", http.StatusOK, resp)
}

// UpdateSaleProduct godoc
// @Router       /v1/sale-products/{id} [put]
// @Summary      Update an existing sale_product
// @Description  Update an existing sale_product with the provided details
// @Tags         sale-products
// @Accept       json
// @Produce      json
// @Param        sale_id       path    string     true    "sale ID to update"
// @Param        sale_product   body    sale_service.UpdateSaleProductRequest true    "Updated data for the sale_product"
// @Success      200  {object}  Response{data=string}
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) UpdateSaleProduct(ctx *gin.Context) {
	var saleProduct = sale_service.SaleProduct{}
	saleProduct.Id = ctx.Param("id")
	if err := ctx.ShouldBindJSON(&saleProduct); err != nil {
		h.handlerResponse(ctx, "error update sale product", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.SaleProductService().Update(ctx.Request.Context(), &sale_service.UpdateSaleProductRequest{
		Id:        saleProduct.Id,
		SaleId:    saleProduct.SaleId,
		ProductId: saleProduct.ProductId,
		Quantity:  saleProduct.Quantity,
		Price:     saleProduct.Price,
	})

	if err != nil {
		h.handlerResponse(ctx, "error sale Update", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "update sale product response", http.StatusOK, resp)
}

// DeleteSaleProduct godoc
// @Router       /v1/sale-products/{id} [delete]
// @Summary      Delete a sale-product
// @Description  delete a sale-product by its unique identifier
// @Tags         sale-products
// @Accept       json
// @Produce      json
// @Param        sale_id   path    string     true    "Sale ID to retrieve"
// @Success      200  {object}  sale_service.Idresponse
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) DeleteSaleProduct(ctx *gin.Context) {
	Id := ctx.Param("id")

	resp, err := h.services.SaleProductService().Delete(ctx.Request.Context(), &sale_service.IdRequest{
		Id: Id,
	})
	if err != nil {
		h.handlerResponse(ctx, "error sale product Delete", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "delete sale product response", http.StatusOK, resp)
}
