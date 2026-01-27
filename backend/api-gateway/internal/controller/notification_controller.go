package controller

import (
	"net/http"

	"api-gateway/internal/dto"
	"api-gateway/internal/helpers/constants"
	notificationpb "api-gateway/internal/pb/notification"
	"api-gateway/internal/utils"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	notificationClient notificationpb.NotificationServiceClient
}

func NewNotificationController(notificationClient notificationpb.NotificationServiceClient) *NotificationController {
	return &NotificationController{
		notificationClient: notificationClient,
	}
}

func (nc *NotificationController) GetNotifications(ctx *gin.Context) {
	var req dto.GetNotificationsRequestDTO
	if err := ctx.ShouldBindQuery(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	userID := ctx.GetString("user_id")
	grpcReq := &notificationpb.GetNotificationsRequest{
		UserId: userID,
		Page:   req.Page,
		Limit:  req.Limit,
	}

	grpcRes, err := nc.notificationClient.GetNotifications(ctx, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}

func (nc *NotificationController) MarkAsRead(ctx *gin.Context) {
	notificationID := ctx.Param("id")
	userID := ctx.GetString("user_id")

	grpcReq := &notificationpb.MarkAsReadRequest{
		NotificationId: notificationID,
		UserId:         userID,
	}

	_, err := nc.notificationClient.MarkAsRead(ctx, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, nil)
	ctx.JSON(http.StatusOK, res)
}

func (nc *NotificationController) SendNotification(ctx *gin.Context) {
	var req dto.SendNotificationRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	grpcReq := &notificationpb.SendNotificationRequest{
		UserId:  req.UserID,
		Title:   req.Title,
		Message: req.Message,
		Type:    req.Type,
	}

	_, err := nc.notificationClient.SendNotification(ctx, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, nil)
	ctx.JSON(http.StatusOK, res)
}
