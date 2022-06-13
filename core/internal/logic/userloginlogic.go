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

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *helper.Result, err error) {
	// todo: add your logic here and delete this line
	loginResponse := new(types.LoginResponse)
	user := new(models.UserBasic)
	ok, err := models.Engine.Where("name=?", req.Name).And("password=?", req.Passwd).Get(user)
	if err != nil {
		fmt.Println("账号密码错误!:", err)
		resp = helper.NewFailResult(helper.FailCode, "账号密码错误!")
		return
	}
	if !ok {
		resp = helper.NewFailResult(helper.FailCode, "账号密码错误!")
		return
	}

	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name, 3000)
	if err != nil {
		fmt.Println("生成token失败:", err)
		resp = helper.NewFailResult(helper.FailCode, "生成token失败!")
		return
	}
	loginResponse.Token = token
	resp = helper.NewSuccessResult(helper.SuccessCode, loginResponse)
	return
}
