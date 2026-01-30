package controller

import (
	"net/http"

	"api-gateway/internal/dto"
	"api-gateway/internal/helper/constant"
	paymentpb "api-gateway/internal/pb/payment"
	"api-gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	paymentClient paymentpb.PaymentServiceClient
}

func NewPaymentController(paymentClient paymentpb.PaymentServiceClient) *PaymentController {
	return &PaymentController{
		paymentClient: paymentClient,
	}
}

func (c *PaymentController) WebhookPayment(ctx *gin.Context) {
	var req dto.MidtransWebhookRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := util.BuildResponseFailed(constant.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	grpcReq := &paymentpb.WebhookRequest{
		OrderId:           req.OrderID,
		TransactionStatus: req.TransactionStatus,
		FraudStatus:       req.FraudStatus,
		PaymentType:       req.PaymentType,
		TransactionId:     req.TransactionID,
	}

	_, err := c.paymentClient.WebhookPayment(ctx, grpcReq)
	if err != nil {
		// Webhook harus tetap return 200 OK ke Payment Gateway agar tidak di-retry terus menerus,
		// meskipun internal error (log error di sisi server).
		// Namun untuk development, return error asli dulu.
		res := util.BuildResponseFailed(constant.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constant.MsgSuccess, nil)
	ctx.JSON(http.StatusOK, res)
}
