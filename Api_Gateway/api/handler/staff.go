package handler

import (
	"net/http"
	"strconv"

	staff_service "api-gateway-service/genproto/staff_service"

	"github.com/gin-gonic/gin"
)

// CreateStaff godoc
// @Router       /v1/staffs [post]
// @Summary      Create a new staff
// @Description  Create a new staff with the provided details
// @Tags         staffs
// @Accept       json
// @Produce      json
// @Param        staff     body  staff_service.CreateStaffRequest true  "data of the staff"
// @Success      201  {object}  staff_service.IdResponse
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) CreateStaff(ctx *gin.Context) {
	var staff = staff_service.Staff{}

	err := ctx.ShouldBindJSON(&staff)
	if err != nil {
		h.handlerResponse(ctx, "CreateStaff", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.StaffService().Create(ctx, &staff_service.CreateStaffRequest{
		BranchId:  staff.BranchId,
		TariffId:  staff.TariffId,
		Name:      staff.Name,
		StaffType: staff.StaffType,
		Login:     staff.Login,
		Password:  staff.Password,
		Phone:     staff.Phone,
	})

	if err != nil {
		h.handlerResponse(ctx, "StaffService().Create", http.StatusBadRequest, err.Error())

		return
	}

	h.handlerResponse(ctx, "create staff response", http.StatusOK, resp)
}

// GetAllStaffs godoc
// @Router       /v1/staffs [get]
// @Summary      GetAll staffs
// @Description  get staffs
// @Tags         staffs
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"  Default(10)
// @Param		 page     query     int  false  "page for response"   Default(1)
// @Param        name     query     string false "search by name"
// @Param        branch_id     query     string false "search by branch_id"
// @Param        tariff_id     query     string false "search by tariff_id"
// @Success      200  {array}   staff_service.Staff
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetAllStaff(ctx *gin.Context) {
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

	balanceFromStr := ctx.Query("balanceFrom")
	balanceFrom, err := strconv.ParseFloat(balanceFromStr, 64)
	if err != nil {
		h.handlerResponse(ctx, "error parsing balanceFrom", http.StatusBadRequest, err.Error())
		return
	}

	balanceToStr := ctx.Query("balanceTo")
	balanceTo, err := strconv.ParseFloat(balanceToStr, 64)
	if err != nil {
		h.handlerResponse(ctx, "error parsing balanceTo", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.StaffService().GetAll(ctx.Request.Context(), &staff_service.GetAllStaffRequest{
		Offset:      int32(page),
		Limit:       int32(limit),
		Search:      ctx.Query("name"),
		BalanceFrom: balanceFrom,
		BalanceTo:   balanceTo,
	})

	if err != nil {
		h.handlerResponse(ctx, "error GetListStaff", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get list staff response", http.StatusOK, resp)
}

// GetStaff godoc
// @Router       /v1/staffs/{id} [get]
// @Summary      Get a staff by ID
// @Description  Retrieve a staff by its unique identifier
// @Tags         staffs
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Staff ID to retrieve"
// @Success      200  {object}  staff_service.Staff
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetStaff(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.StaffService().Get(ctx.Request.Context(), &staff_service.IdRequest{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error staff GetById", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get staff response", http.StatusOK, resp)
}

// UpdateStaff godoc
// @Router       /v1/staffs/{id} [put]
// @Summary      Update an existing staff
// @Description  Update an existing staff with the provided details
// @Tags         staffs
// @Accept       json
// @Produce      json
// @Param        id       path    string     true    "Staff ID to update"
// @Param        staff   body    staff_service.UpdateStaffRequest  true    "Updated data for the staff"
// @Success      200  {object}  Response{data=string}
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) UpdateStaff(ctx *gin.Context) {
	var staff = staff_service.Staff{}
	staff.Id = ctx.Param("id")
	err := ctx.ShouldBindJSON(&staff)
	if err != nil {
		h.handlerResponse(ctx, "error while binding", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.StaffService().Update(ctx.Request.Context(), &staff_service.UpdateStaffRequest{
		Id:        staff.Id,
		BranchId:  staff.BranchId,
		TariffId:  staff.TariffId,
		StaffType: staff.StaffType,
		Name:      staff.Name,
		Balance:   staff.Balance,
		Login:     staff.Login,
		Password:  staff.Password,
		Phone:     staff.Phone,
	})

	if err != nil {
		h.handlerResponse(ctx, "error staff Update", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "update staff response", http.StatusOK, resp)
}

// DeleteStaff godoc
// @Router       /v1/staffs/{id} [delete]
// @Summary      Delete a staff
// @Description  delete a staff by its unique identifier
// @Tags         staffs
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Staff ID to retrieve"
// @Success      200  {object}  staff_service.IdRequest
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) DeleteStaff(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.StaffService().Delete(ctx.Request.Context(), &staff_service.IdRequest{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error staff Delete", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "delete staff response", http.StatusOK, resp)
}
