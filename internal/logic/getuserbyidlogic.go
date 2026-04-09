// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"log"

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
	existingUser, _ := l.svcCtx.UserMod.GetUserByID(l.ctx, req.Id)

	log.Println("User id: ", req.Id)

	if existingUser == nil {
		panic(xerr.NewException(404, "Người dùng không tồn tại"))
	}

	user := &types.ResponseGetUserByIdData{
		Id:        existingUser.ID,
		Username:  existingUser.Username,
		Name:      existingUser.Name,
		CreatedAt: existingUser.CreatedAt.String(),
	}

	return &types.ResponseGetUserById{
		Code:    0,
		Message: "Thành công",
		Data:    user,
	}, nil
}
