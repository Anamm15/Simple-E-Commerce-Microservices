package controller

import (
	"net/http"

	"api-gateway/internal/dto"
	"api-gateway/internal/helper/constant"
	inventorypb "api-gateway/internal/pb/inventory"
	"api-gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	InventoryClient inventorypb.InventoryServiceClient
}

func NewInventoryController(inventoryClient inventorypb.InventoryServiceClient) *InventoryController {
	return &InventoryController{
		InventoryClient: inventoryClient,
	}
}

func (c *InventoryController) UpdateStock(ctx *gin.Context) {
	var req dto.UpdateStockRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := util.BuildResponseFailed(constant.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	grpcReq := &inventorypb.UpdateStockRequest{
		ProductId:  req.ProductID,
		UpdateType: req.UpdateType,
		Quantity:   req.Quantity,
	}

	grpcRes, err := c.InventoryClient.UpdateStock(ctx, grpcReq)
	if err != nil {
		res := util.BuildResponseFailed(constant.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constant.MsgSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}
