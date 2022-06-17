package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.Result, err error) {
	// todo: add your logic here and delete this line
	uc := new(models.UserBasic)
	count, err := models.Engine.Where("name=?", req.Name).Count(uc)
	if err != nil {
		l.Logger.Errorf("查询错误:%v", err)
		resp = helper.NewFailResult(helper.FailCode, err.Error())
		return
	}
	if count > 0 {
		resp = helper.NewFailResult(helper.FailDoubleCode, "数据重复!")
		return
	}
	// 数据入库
	user := &models.UserBasic{
		Identity: helper.UUID(),
		Name:     req.Name,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
	}

	_, err = models.Engine.Insert(user)
	if err != nil {
		resp = helper.NewFailResult(helper.FailCode, "报错失败!")
		return
	}
	resp = helper.NewSuccessResult(helper.SuccessCode, nil)
	return
}
