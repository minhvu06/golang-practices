// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"time"

	"user-api/internal/svc"
	"user-api/internal/types"
	"user-api/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.RequestGetUserById) (resp *types.ResponseGetUserById, err error) {
	if req.Id == "12345" {
		panic(xerr.NewException(1001, "User not found"))
	}

	user := &types.ResponseGetUserByIdData{
		Id:        req.Id,
		Username:  "User",
		Name:      "Người dùng",
		CreatedAt: time.Now().UTC().String(),
	}

	return &types.ResponseGetUserById{
		Code:    0,
		Message: "Thành công",
		Data:    user,
	}, nil
}
