// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"time"

	"user-api/internal/model"
	"user-api/internal/svc"
	"user-api/internal/types"
	"user-api/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
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
	existingUser, _ := l.svcCtx.UserMod.GetUserByUsername(l.ctx, req.Username)
	if existingUser != nil {
		panic(xerr.NewException(400, "Tên tài khoản đã được sử dụng"))
	}

	hashedPwd, err := hashPassword(req.Password)
	if err != nil {
		panic(xerr.NewException(409, "Lỗi xử lý hệ thống"))
	}

	user := &model.User{
		Username:  req.Username,
		Name:      req.Name,
		Password:  hashedPwd,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := l.svcCtx.UserMod.CreateUser(l.ctx, user); err != nil {
		panic(xerr.NewException(409, "Không thể tạo người dùng"))
	}

	return &types.ResponseCreateUser{
		Code:    0,
		Message: "Thành công",
	}, nil
}

func hashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash), err
}
