package controller

import (
	"net/http"

	"api-gateway/internal/dto"
	"api-gateway/internal/helpers/constants"
	authpb "api-gateway/internal/pb/auth"
	"api-gateway/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserClient authpb.AuthServiceClient
}

func NewAuthController(userClient authpb.AuthServiceClient) *AuthController {
	return &AuthController{
		UserClient: userClient,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req dto.RegisterRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	grpcReq := &authpb.RegisterRequest{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}

	grpcRes, err := c.UserClient.Register(ctx, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgRegisterSuccess, grpcRes)
	ctx.JSON(http.StatusCreated, res)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req dto.LoginRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	grpcReq := &authpb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	grpcRes, err := c.UserClient.Login(ctx, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInvalidCredentials, err.Error(), nil)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	ctx.SetCookie("refresh_token", grpcRes.RefreshToken, 7*24*3600, "/", "", false, true)
	res := utils.BuildResponseSuccess(constants.MsgLoginSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}

func (c *AuthController) RefreshToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		var req dto.RefreshTokenRequestDTO
		if bindErr := ctx.ShouldBindJSON(&req); bindErr == nil {
			refreshToken = req.RefreshToken
		}
	}

	if refreshToken == "" {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, "Refresh token is missing", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	grpcReq := &authpb.RefreshTokenRequest{
		RefreshToken: refreshToken,
	}

	grpcRes, err := c.UserClient.RefreshToken(ctx, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	if grpcRes.RefreshToken != "" {
		ctx.SetCookie("refresh_token", grpcRes.RefreshToken, 7*24*3600, "/", "", false, true)
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, grpcRes)
	ctx.JSON(http.StatusOK, res)
}

func (c *AuthController) Logout(ctx *gin.Context) {
	ctx.SetCookie("refresh_token", "", -1, "/", "", false, true)

	res := utils.BuildResponseSuccess(constants.MsgLogoutSuccess, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *AuthController) ChangePassword(ctx *gin.Context) {
	var req dto.ChangePasswordRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	userID := ctx.GetString("user_id")
	if userID == "" {
		res := utils.BuildResponseFailed(constants.MsgUnauthorized, "User ID not found in context", nil)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	grpcReq := &authpb.ChangePasswordRequest{
		UserId:      userID,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}

	_, err := c.UserClient.ChangePassword(ctx, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *AuthController) ResetPassword(ctx *gin.Context) {
	var req dto.ResetPasswordRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MsgInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	grpcReq := &authpb.ResetPasswordRequest{
		ResetToken:  req.ResetToken,
		NewPassword: req.NewPassword,
	}

	_, err := c.UserClient.ResetPassword(ctx, grpcReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MsgInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MsgSuccess, nil)
	ctx.JSON(http.StatusOK, res)
}
