package controller

import (
	"net/http"

	"api-gateway/internal/dto"
	"api-gateway/internal/helpers/constants"
	userpb "api-gateway/internal/pb/user"
	"api-gateway/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userClient userpb.UserServiceClient
}

func NewUserController(userClient userpb.UserServiceClient) *UserController {
	return &UserController{
		userClient: userClient,
	}
}

func (uc *UserController) GetUserProfile(c *gin.Context) {
	userID := c.GetString("user_id")

	grpcReq := &userpb.GetUserProfileRequest{
		UserId: userID,
	}

	grpcRes, err := uc.userClient.GetUserProfile(c, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, grpcRes)
	c.JSON(http.StatusOK, res)
}

func (uc *UserController) UpdateUserProfile(c *gin.Context) {
	var req dto.UpdateUserProfileRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	userID := c.GetString("user_id")

	grpcReq := &userpb.UpdateUserProfileRequest{
		UserId:      userID,
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
	}

	grpcRes, err := uc.userClient.UpdateUserProfile(c, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgUserUpdated, grpcRes)
	c.JSON(http.StatusOK, res)
}

func (uc *UserController) AddAddress(c *gin.Context) {
	var req dto.AddAddressRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	userID := c.GetString("user_id")

	grpcReq := &userpb.AddAddressRequest{
		UserId:    userID,
		Street:    req.Street,
		City:      req.City,
		Province:  req.Province,
		ZipCode:   req.ZipCode,
		IsPrimary: req.IsPrimary,
	}

	grpcRes, err := uc.userClient.AddAddress(c, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgAddressAdded, grpcRes)
	c.JSON(http.StatusCreated, res)
}

func (uc *UserController) GetAddress(c *gin.Context) {
	addressID := c.Param("id")
	userID := c.GetString("user_id")

	grpcReq := &userpb.GetAddressRequest{
		AddressId: addressID,
		UserId:    userID,
	}

	grpcRes, err := uc.userClient.GetAddress(c, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, grpcRes)
	c.JSON(http.StatusOK, res)
}

func (uc *UserController) UpdateAddress(c *gin.Context) {
	addressID := c.Param("id")
	userID := c.GetString("user_id")

	var req dto.UpdateAddressRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	grpcReq := &userpb.UpdateAddressRequest{
		AddressId: addressID,
		UserId:    userID,
		Street:    req.Street,
		City:      req.City,
		Province:  req.Province,
		ZipCode:   req.ZipCode,
		IsPrimary: req.IsPrimary,
	}

	grpcRes, err := uc.userClient.UpdateAddress(c, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgAddressUpdated, grpcRes)
	c.JSON(http.StatusOK, res)
}

func (uc *UserController) DeleteAddress(c *gin.Context) {
	addressID := c.Param("id")
	userID := c.GetString("user_id")

	grpcReq := &userpb.DeleteAddressRequest{
		AddressId: addressID,
		UserId:    userID,
	}

	_, err := uc.userClient.DeleteAddress(c, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgAddressDeleted, nil)
	c.JSON(http.StatusOK, res)
}
