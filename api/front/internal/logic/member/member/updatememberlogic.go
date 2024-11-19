package member

import (
	"context"
	"errors"
	"github.com/summmer-gonner/traffica/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/summmer-gonner/traffica/api/front/internal/svc"
	"github.com/summmer-gonner/traffica/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberLogic 更新会员信息
/*
Author: LiuFeiHua
Date: 2024/5/16 14:07
*/
type UpdateMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberLogic {
	return &UpdateMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMember 更新会员信息
func (l *UpdateMemberLogic) UpdateMember(req *types.UpdateMemberReq) (resp *types.UpdateMemberResp, err error) {
	_, err = l.svcCtx.MemberService.UpdateMember(l.ctx, &umsclient.UpdateMemberReq{
		Id:                    req.Id,
		MemberLevelId:         req.MemberLevelId,
		MemberName:            req.Username,
		Phone:                 req.Phone,
		MemberStatus:          req.Status,
		Icon:                  req.Icon,
		Gender:                req.Gender,
		Birthday:              req.Birthday,
		City:                  req.City,
		Job:                   req.Job,
		PersonalizedSignature: req.PersonalizedSignature,
		SourceType:            req.SourceType,
		Integration:           req.Integration,
		Growth:                req.Growth,
		LotteryCount:          req.LuckeyCount,
		HistoryIntegration:    req.HistoryIntegration,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员信息失败,参数:%+v,异常:%s", req, err.Error())
		return nil, errors.New("更新会员信息失败")
	}

	return &types.UpdateMemberResp{
		Code:    "000000",
		Message: "更新会员信息成功",
	}, nil
}
