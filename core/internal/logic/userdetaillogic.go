package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"
	"fmt"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.DetailRequest) (resp *helper.Result, err error) {
	// todo: add your logic here and delete this line
	user := new(models.UserBasic)
	ok, err := models.Engine.Where("identity=?", req.Identity).Get(user)
	if err != nil {
		fmt.Println("错误:", err)
		resp = helper.NewFailResult(helper.FailCode, "异常")
		return
	}
	if !ok {
		resp = helper.NewFailResult(helper.FailCode, "identity错误!")
		return
	}

	resp = helper.NewSuccessResult(helper.SuccessCode, types.DetailResponse{
		Identity: user.Identity,
		Name:     user.Name,
		Email:    user.Email,
	})

	return
}
