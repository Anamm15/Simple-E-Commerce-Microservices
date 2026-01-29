package controller

import (
	"encoding/json"
	"net/http"

	"api-gateway/internal/dto"
	"api-gateway/internal/helper/constants"
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
		res := util.BuildResponseFailed(constants.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Marshal struct ke JSON string untuk disimpan sebagai raw log audit trail
	rawJSON, _ := json.Marshal(req)

	grpcReq := &paymentpb.WebhookRequest{
		OrderId:           req.OrderID,
		TransactionStatus: req.TransactionStatus,
		FraudStatus:       req.FraudStatus,
		PaymentType:       req.PaymentType,
		TransactionId:     req.TransactionID,
		RawJsonPayload:    string(rawJSON),
	}

	_, err := c.paymentClient.WebhookPayment(ctx, grpcReq)
	if err != nil {
		// Webhook harus tetap return 200 OK ke Payment Gateway agar tidak di-retry terus menerus,
		// meskipun internal error (log error di sisi server).
		// Namun untuk development, return error asli dulu.
		res := util.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := util.BuildResponseSuccess(constants.MsgSuccess, nil)
	ctx.JSON(http.StatusOK, res)
}
