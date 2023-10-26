package handler

import (
	"net/http"
	"strconv"

	staff_service "api-gateway-service/genproto/staff_service"

	"github.com/gin-gonic/gin"
)

// CreateStaffTariff godoc
// @Router       /v1/staff-tariffs [post]
// @Summary      Create a new staff tariff
// @Description  Create a new staff tariff with the provided details
// @Tags         staff-tariffs
// @Accept       json
// @Produce      json
// @Param        tariff     body  staff_service.CreateStaffTarifRequest true  "data of the staff tariff"
// @Success      201  {object}  staff_service.IdResponse
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) CreateStaffTariff(ctx *gin.Context) {
	var staffTariff = staff_service.StaffTarif{}

	err := ctx.ShouldBindJSON(&staffTariff)
	if err != nil {
		h.handlerResponse(ctx, "CreateStaffTariff", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.StaffTariffService().Create(ctx, &staff_service.CreateStaffTarifRequest{
		Name:          staffTariff.Name,
		Type:          staffTariff.Type,
		AmountForCash: staffTariff.AmountForCash,
		AmountForCard: staffTariff.AmountForCard,
	})

	if err != nil {
		h.handlerResponse(ctx, "StaffTariffService().Create", http.StatusBadRequest, err.Error())

		return
	}

	h.handlerResponse(ctx, "create staff tariff response", http.StatusOK, resp)
}

// ListStaffTariff godoc
// @Router       /v1/staff-tariffs [get]
// @Summary      List staff-tariffs
// @Description  get staff-tariffs
// @Tags         staff-tariffs
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"  Default(10)
// @Param		 page     query     int  false  "page for response"   Default(1)
// @Param        type     query     string false "search by type"
// @Success      200  {array}   staff_service.StaffTariff
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetAllStaffTariff(ctx *gin.Context) {
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

	resp, err := h.services.StaffTariffService().GetAll(ctx.Request.Context(), &staff_service.GetAllStaffTarifRequest{
		Offset: int32(page),
		Limit:  int32(limit),
		Search: ctx.Query("type"),
	})

	if err != nil {
		h.handlerResponse(ctx, "error GetAllStaffTariff", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get all staff tariff response", http.StatusOK, resp)
}

// GetStaffTariff godoc
// @Router       /v1/staff-tariffs/{id} [get]
// @Summary      Get a staff-tariff by ID
// @Description  Retrieve a staff-tariff by its unique identifier
// @Tags         staff-tariffs
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Staff Tariff ID to retrieve"
// @Success      200  {object}  staff_service.StaffTariff
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetStaffTariff(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.StaffTariffService().Get(ctx.Request.Context(), &staff_service.IdRequest{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error staff tariff GetById", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get staff tariff response", http.StatusOK, resp)
}

// UpdateStaffTariff godoc
// @Router       /v1/staff-tariffs/{id} [put]
// @Summary      Update an existing staff-tariff
// @Description  Update an existing staff-tariff with the provided details
// @Tags         staff-tariffs
// @Accept       json
// @Produce      json
// @Param        id       path    string     true    "Staff Tariff ID to update"
// @Param        staff-tariff   body    staff_service.UpdateStaffTarifRequest  true    "Updated data for the staff-tariff"
// @Success      200  {object}   Response{data=string}
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) UpdateStaffTariff(ctx *gin.Context) {
	var staffTariff = staff_service.StaffTarif{}
	staffTariff.Id = ctx.Param("id")
	err := ctx.ShouldBindJSON(&staffTariff)
	if err != nil {
		h.handlerResponse(ctx, "error while binding", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.StaffTariffService().Update(ctx.Request.Context(), &staff_service.UpdateStaffTarifRequest{
		Id:            staffTariff.Id,
		Name:          staffTariff.Name,
		Type:          staffTariff.Type,
		AmountForCash: staffTariff.AmountForCash,
		AmountForCard: staffTariff.AmountForCard,
	})

	if err != nil {
		h.handlerResponse(ctx, "error staff tariff Update", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "update staff tariff response", http.StatusOK, resp)
}

// DeleteStaffTariff godoc
// @Router       /v1/staff-tariffs/{id} [delete]
// @Summary      Delete a staff-tariff
// @Description  delete a staff-tariff by its unique identifier
// @Tags         staff-tariffs
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "StaffTariff ID to retrieve"
// @Success      200  {object}  staff_service.IdRequest
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) DeleteStaffTariff(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.StaffTariffService().Delete(ctx.Request.Context(), &staff_service.IdRequest{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error staff tariff Delete", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "delete staff tariff response", http.StatusOK, resp)
}
