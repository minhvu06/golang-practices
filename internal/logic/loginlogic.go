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
	"golang.org/x/crypto/bcrypt"
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
	existingUser, _ := l.svcCtx.UserMod.GetUserByUsername(l.ctx, req.Username)

	if existingUser == nil {
		panic(xerr.NewException(400, "Tài khoản hoặc mật khẩu không chính xác"))
	}

	if !verifyPassword(existingUser.Password, req.Password) {
		panic(xerr.NewException(400, "Tài khoản hoặc mật khẩu không chính xác"))
	}

	accessToken, _ := helper.GenerateToken(existingUser.Username)

	res := &types.ResponseLoginData{
		Id:          existingUser.ID,
		Username:    existingUser.Username,
		Name:        existingUser.Name,
		AccessToken: accessToken,
	}

	return &types.ResponseLogin{
		Code:    0,
		Message: "Thành công",
		Data:    res,
	}, nil
}

func verifyPassword(hash, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
