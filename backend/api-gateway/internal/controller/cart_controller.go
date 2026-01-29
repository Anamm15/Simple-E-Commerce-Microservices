package controller

import (
	"net/http"

	"api-gateway/internal/dto"
	"api-gateway/internal/helper/constant"
	cartpb "api-gateway/internal/pb/cart"
	"api-gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	CartClient cartpb.CartServiceClient
}

func NewCartController(cartClient cartpb.CartServiceClient) *CartController {
	return &CartController{
		CartClient: cartClient,
	}
}

func (c *CartController) AddItem(ctx *gin.Context) {
	var req dto.AddCartItemRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := util.BuildResponseFailed(constant.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	userID := ctx.GetString("user_id")
	grpcReq := &cartpb.AddItemRequest{
		UserId:    userID,
		ProductId: req.ProductID,
		Quantity:  req.Quantity,
	}

	grpcRes, err := c.CartClient.AddItem(ctx, grpcReq)
	if err != nil {
		res := util.BuildResponseFailed(constant.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constant.MsgSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}

func (c *CartController) GetCart(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	grpcReq := &cartpb.GetCartRequest{
		UserId: userID,
	}

	grpcRes, err := c.CartClient.GetCart(ctx, grpcReq)
	if err != nil {
		res := util.BuildResponseFailed(constant.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constant.MsgSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}

func (c *CartController) RemoveItem(ctx *gin.Context) {
	productID := ctx.Param("product_id")
	userID := ctx.GetString("user_id")

	grpcReq := &cartpb.RemoveItemRequest{
		UserId:    userID,
		ProductId: productID,
	}

	grpcRes, err := c.CartClient.RemoveItem(ctx, grpcReq)
	if err != nil {
		res := util.BuildResponseFailed(constant.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constant.MsgSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}

func (c *CartController) ClearCart(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	grpcReq := &cartpb.ClearCartRequest{
		UserId: userID,
	}

	_, err := c.CartClient.ClearCart(ctx, grpcReq)
	if err != nil {
		res := util.BuildResponseFailed(constant.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constant.MsgSuccess, nil)
	ctx.JSON(http.StatusOK, res)
}
