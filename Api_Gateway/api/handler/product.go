package handler

import (
	"net/http"
	"strconv"

	product_service "api-gateway-service/genproto/product_service"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Router       /v1/products [post]
// @Summary      Create a new product
// @Description  Create a new product with the provided details
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product     body  product_service.CreateProductRequest true  "data of the product"
// @Success      201  {object}  product_service.IdResponse
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) CreateProduct(ctx *gin.Context) {
	var product = product_service.Product{}

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		h.handlerResponse(ctx, "CreateProduct", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.ProductService().Create(ctx, &product_service.CreateProductRequest{
		Name:       product.Name,
		Price:      product.Price,
		Barcode:    product.Barcode,
		CategoryId: product.CategoryId,
	})

	if err != nil {
		h.handlerResponse(ctx, "ProductService().Create", http.StatusBadRequest, err.Error())

		return
	}

	h.handlerResponse(ctx, "create product response", http.StatusOK, resp)
}

// GetAllProducts godoc
// @Router       /v1/products [get]
// @Summary      GetAll products
// @Description  get products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"  Default(10)
// @Param		 page     query     int  false  "page for response"   Default(1)
// @Param        name     query     string false "search by name"
// @Param        barcode     query     string false "search by barcode"
// @Success      200  {array}   product_service.Product
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetAllProduct(ctx *gin.Context) {
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

	resp, err := h.services.ProductService().GetAll(ctx.Request.Context(), &product_service.GetAllProductRequest{
		Offset: int32(page),
		Limit:  int32(limit),
		Search: ctx.Query("search"),
	})

	if err != nil {
		h.handlerResponse(ctx, "error GetAllProduct", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get All product response", http.StatusOK, resp)
}

// GetProduct godoc
// @Router       /v1/products/{id} [get]
// @Summary      Get a product by ID
// @Description  Retrieve a product by its unique identifier
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Product ID to retrieve"
// @Success      200  {object}  product_service.Product
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.ProductService().Get(ctx.Request.Context(), &product_service.IdRequest{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error product GetById", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get product response", http.StatusOK, resp)
}

// UpdateProduct godoc
// @Router       /v1/products/{id} [put]
// @Summary      Update an existing product
// @Description  Update an existing product with the provided details
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       path    string     true    "Product ID to update"
// @Param        product   body    product_service.UpdateProductRequest  true    "Updated data for the product"
// @Success      200  {object}  Response{data=string}
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) UpdateProduct(ctx *gin.Context) {
	var product = product_service.Product{}
	product.Id = ctx.Param("id")
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		h.handlerResponse(ctx, "error while binding", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.ProductService().Update(ctx.Request.Context(), &product_service.UpdateProductRequest{
		Id:         product.Id,
		Name:       product.Name,
		Price:      product.Price,
		Barcode:    product.Barcode,
		CategoryId: product.CategoryId,
	})

	if err != nil {
		h.handlerResponse(ctx, "error product Update", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "update product response", http.StatusOK, resp)
}

// DeleteProduct godoc
// @Router       /v1/products/{id} [delete]
// @Summary      Delete a product
// @Description  delete a product by its unique identifier
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Product ID to retrieve"
// @Success      200  {object}  product_service.IdResponse
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.ProductService().Delete(ctx.Request.Context(), &product_service.IdRequest{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error product Delete", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "delete product response", http.StatusOK, resp)
}
