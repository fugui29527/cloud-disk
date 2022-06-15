package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/types"
	"context"

	"cloud-disk/core/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthorizationLogic {
	return &RefreshAuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAuthorizationLogic) RefreshAuthorization(token string) (resp *types.Result, err error) {
	uc, err := helper.AnalyzeToken(token)
	if err != nil {
		resp = helper.NewFailResult(helper.FailAuthCode, err.Error())
		l.Logger.Errorf("解析token错误:%v", err)
		return
	}
	// 根据 UserClaim 中的信息，生成新的 Token
	nemToken, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, define.TokenExpire)
	if err != nil {
		resp = helper.NewFailResult(helper.FailCode, err.Error())
		l.Logger.Errorf("生成token错误:%v", err)
		return
	}
	// 生成新的 Refresh Token
	newRefreshToken, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, define.RefreshTokenExpire)
	if err != nil {
		resp = helper.NewFailResult(helper.FailCode, err.Error())
		l.Logger.Errorf("解析token错误:%v", err)
		return
	}
	refreshAuthorizationReply := new(types.RefreshAuthorizationReply)
	refreshAuthorizationReply.Token = nemToken
	refreshAuthorizationReply.RefreshToken = newRefreshToken
	resp = helper.NewSuccessResult(helper.SuccessCode, refreshAuthorizationReply)
	return
}
