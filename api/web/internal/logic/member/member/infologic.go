package member

import (
	"context"
	"encoding/json"
	"github.com/summmer-gonner/traffica/rpc/sms/smsclient"
	"github.com/summmer-gonner/traffica/rpc/ums/umsclient"

	"github.com/summmer-gonner/traffica/api/web/internal/svc"
	"github.com/summmer-gonner/traffica/api/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// InfoLogic
/*
Author: LiuFeiHua
Date: 2024/4/7 18:07
*/
type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Info 获取个人信息
func (l *InfoLogic) Info() (*types.InfoResp, error) {
	//从jwt中提取会员id
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	member, _ := l.svcCtx.MemberService.QueryMemberDetail(l.ctx, &umsclient.QueryMemberDetailReq{Id: memberId})

	// 获取用户优惠券
	result, _ := l.svcCtx.CouponHistoryService.QueryCouponCount(l.ctx, &smsclient.QueryCouponCountReq{
		MemberId: memberId,
	})

	var count int32 = 0
	//获取不到优惠券数量的时候,默认返回0
	if result != nil {
		count = int32(result.Total)
	}

	data := types.MemberData{
		Id:                    member.Id,
		MemberLevelId:         member.MemberLevelId,
		Username:              member.MemberName,
		Nickname:              member.Nickname,
		Phone:                 member.Phone,
		Status:                member.MemberStatus,
		CreateTime:            member.CreateTime,
		Icon:                  member.Icon,
		Gender:                member.Gender,
		Birthday:              member.Birthday,
		City:                  member.City,
		Job:                   member.Job,
		PersonalizedSignature: member.PersonalizedSignature,
		SourceType:            member.SourceType,
		Integration:           member.Integration,
		Growth:                member.Growth,
		LuckeyCount:           member.LotteryCount,
		HistoryIntegration:    member.HistoryIntegration,
		CouponCount:           count,
	}
	return &types.InfoResp{
		Code:    0,
		Message: "查询会员信息",
		Data:    data,
	}, nil
}
