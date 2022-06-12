package logic

import (
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoreLogic {
	return &CoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoreLogic) Core(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)
	user := make([]*models.UserBasic, 0)
	err = models.Engine.Find(&user)
	if err != nil {
		fmt.Println("===============", err)
		resp.Message = err.Error()
		return
	}
	fmt.Println(user)
	marshal, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		resp.Message = err.Error()
		return
	}
	resp.Message = string(marshal)
	return
}
