// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"user-api/internal/svc"
	"user-api/internal/types"
	"user-api/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.RequestCreateUser) (resp *types.ResponseCreateUser, err error) {
	if req.Username == "vnm" {
		panic(xerr.NewException(400, "Tên tài khoản đã được sử dụng"))
	}

	return &types.ResponseCreateUser{
		Code:    0,
		Message: "Thành công",
	}, nil
}
