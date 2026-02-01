package controller

import (
	"net/http"

	"api-gateway/internal/dto"
	"api-gateway/internal/helper/constant"
	shippingpb "api-gateway/internal/pb/shipping"
	"api-gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

type ShippingController struct {
	shippingClient shippingpb.ShippingServiceClient
}

func NewShippingController(shippingClient shippingpb.ShippingServiceClient) *ShippingController {
	return &ShippingController{
		shippingClient: shippingClient,
	}
}

func (c *ShippingController) OfferShippingCost(ctx *gin.Context) {
	var req dto.OfferShippingCostRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := util.BuildResponseFailed(constant.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	grpcReq := &shippingpb.ShippingCostOfferRequest{
		OriginCity:      req.OriginCity,
		DestinationCity: req.DestinationCity,
		TotalWeightG:    req.TotalWeightG,
	}

	grpcRes, err := c.shippingClient.OfferShippingCost(ctx, grpcReq)
	if err != nil {
		res := util.BuildResponseFailed(constant.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constant.MsgSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}
