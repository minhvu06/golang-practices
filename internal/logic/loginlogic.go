// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"user-api/internal/helper"
	"user-api/internal/svc"
	"user-api/internal/types"
	"user-api/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.RequestLogin) (resp *types.ResponseLogin, err error) {
	if req.Username != "vnm" || req.Password != "123456" {
		panic(xerr.NewException(400, "Tài khoản hoặc mật khẩu không chính xác"))
	}

	accessToken, _ := helper.GenerateToken(req.Username)

	res := &types.ResponseLoginData{
		Id:          req.Username,
		Username:    req.Username,
		Name:        req.Username,
		AccessToken: accessToken,
	}

	return &types.ResponseLogin{
		Code:    0,
		Message: "Thành công",
		Data:    res,
	}, nil
}
