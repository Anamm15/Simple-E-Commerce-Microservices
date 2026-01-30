package controller

import (
	"net/http"

	"api-gateway/internal/dto"
	"api-gateway/internal/helper/constant"
	orderpb "api-gateway/internal/pb/order"
	"api-gateway/pkg/util"

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
		res := util.BuildResponseFailed(constant.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	userID := ctx.MustGet("user_id").(string)
	grpcReq := &orderpb.CheckoutRequest{
		UserId:      userID,
		Address:     req.Address,
		CourierCode: req.CourierCode,
		ServiceCode: req.ServiceCode,
		ProductIds:  req.ProdcutIDs,
	}

	grpcRes, err := c.orderService.Checkout(ctx, grpcReq)
	if err != nil {
		res := util.BuildResponseFailed(constant.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constant.MsgSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}

func (c *OrderController) GetOrderHistory(ctx *gin.Context) {
	var req dto.GetOrderHistoryRequestDTO
	if err := ctx.ShouldBindQuery(&req); err != nil {
		res := util.BuildResponseFailed(constant.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Mapping string status to Enum int32
	statusEnum := orderpb.OrderStatus(orderpb.OrderStatus_value[req.StatusFilter])

	userID := ctx.MustGet("user_id").(string)
	grpcReq := &orderpb.GetOrderHistoryRequest{
		UserId:       userID,
		Page:         req.Page,
		Limit:        req.Limit,
		StatusFilter: statusEnum,
	}

	grpcRes, err := c.orderService.GetOrderHistory(ctx, grpcReq)
	if err != nil {
		res := util.BuildResponseFailed(constant.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constant.MsgSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}

func (c *OrderController) GetOrderDetail(ctx *gin.Context) {
	orderID := ctx.Param("id")
	userID := ctx.MustGet("user_id").(string)

	grpcReq := &orderpb.GetOrderDetailRequest{
		OrderId: orderID,
		UserId:  userID,
	}

	grpcRes, err := c.orderService.GetOrderDetail(ctx, grpcReq)
	if err != nil {
		res := util.BuildResponseFailed(constant.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constant.MsgSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}

func (c *OrderController) GetAllOrders(ctx *gin.Context) {
	var req dto.AdminOrderFilterRequestDTO
	if err := ctx.ShouldBindQuery(&req); err != nil {
		res := util.BuildResponseFailed(constant.MsgInvalidRequest, err.Error(), nil)
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
		res := util.BuildResponseFailed(constant.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constant.MsgSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}

func (c *OrderController) UpdateOrderStatus(ctx *gin.Context) {
	orderID := ctx.Param("id")

	var req dto.UpdateOrderStatusRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := util.BuildResponseFailed(constant.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Validate status validity
	statusVal, ok := orderpb.OrderStatus_value[req.NewStatus]
	if !ok {
		res := util.BuildResponseFailed(constant.MsgInvalidRequest, "Invalid status", nil)
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
		res := util.BuildResponseFailed(constant.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constant.MsgSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}
