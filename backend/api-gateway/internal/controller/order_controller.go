package controller

import (
	"net/http"

	"api-gateway/internal/dto"
	"api-gateway/internal/helpers/constants"
	orderpb "api-gateway/internal/pb/order"
	"api-gateway/internal/utils"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService orderpb.OrderServiceClient
}

func NewOrderController(orderService orderpb.OrderServiceClient) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (c *OrderController) Checkout(ctx *gin.Context) {
	var req dto.CheckoutRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	userID := ctx.GetString("user_id")
	grpcReq := &orderpb.CheckoutRequest{
		UserId:        userID,
		AddressId:     req.AddressID,
		CourierCode:   req.CourierCode,
		ServiceCode:   req.ServiceCode,
		PaymentMethod: req.PaymentMethod,
	}

	grpcRes, err := c.orderService.Checkout(ctx, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}

func (c *OrderController) GetOrderHistory(ctx *gin.Context) {
	var req dto.GetOrderHistoryRequestDTO
	if err := ctx.ShouldBindQuery(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Mapping string status to Enum int32
	statusEnum := orderpb.OrderStatus(orderpb.OrderStatus_value[req.StatusFilter])

	userID := ctx.GetString("user_id")
	grpcReq := &orderpb.GetOrderHistoryRequest{
		UserId:       userID,
		Page:         req.Page,
		Limit:        req.Limit,
		StatusFilter: statusEnum,
	}

	grpcRes, err := c.orderService.GetOrderHistory(ctx, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}

func (c *OrderController) GetOrderDetail(ctx *gin.Context) {
	orderID := ctx.Param("id")
	userID := ctx.GetString("user_id")

	grpcReq := &orderpb.GetOrderDetailRequest{
		OrderId: orderID,
		UserId:  userID,
	}

	grpcRes, err := c.orderService.GetOrderDetail(ctx, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}

func (c *OrderController) GetAllOrders(ctx *gin.Context) {
	var req dto.AdminOrderFilterRequestDTO
	if err := ctx.ShouldBindQuery(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	statusEnum := orderpb.OrderStatus(orderpb.OrderStatus_value[req.Status])

	grpcReq := &orderpb.AdminFilterRequest{
		Page:      req.Page,
		Limit:     req.Limit,
		UserId:    req.UserID,
		Status:    statusEnum,
		DateStart: req.DateStart,
		DateEnd:   req.DateEnd,
	}

	grpcRes, err := c.orderService.GetAllOrders(ctx, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}

func (c *OrderController) UpdateOrderStatus(ctx *gin.Context) {
	orderID := ctx.Param("id")

	var req dto.UpdateOrderStatusRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Validate status validity
	statusVal, ok := orderpb.OrderStatus_value[req.NewStatus]
	if !ok {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, "Invalid status", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	grpcReq := &orderpb.UpdateStatusRequest{
		OrderId:   orderID,
		NewStatus: orderpb.OrderStatus(statusVal),
		Notes:     req.Notes,
	}

	grpcRes, err := c.orderService.UpdateOrderStatus(ctx, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}
