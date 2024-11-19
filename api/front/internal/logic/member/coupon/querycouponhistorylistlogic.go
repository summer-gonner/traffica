package coupon

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCouponHistoryListLogic 获取会员优惠券历史列表
/*
Author: LiuFeiHua
Date: 2024/5/16 16:20
*/
type QueryCouponHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCouponHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponHistoryListLogic {
	return &QueryCouponHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryCouponHistoryList 获取会员优惠券历史列表
func (l *QueryCouponHistoryListLogic) QueryCouponHistoryList(req *types.ListCouponHistoryReq) (resp *types.ListCouponHistoryResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	historyList, err := l.svcCtx.CouponHistoryService.QueryCouponHistoryList(l.ctx, &smsclient.QueryCouponHistoryListReq{
		PageNum:   1,
		PageSize:  100,
		MemberId:  memberId,
		UseStatus: req.UseStatus,
	})
	if err != nil {
		return nil, err
	}
	var list []*types.ListCouponHistoryData

	for _, item := range historyList.List {
		list = append(list, &types.ListCouponHistoryData{
			Id:             item.Id,
			CouponId:       item.CouponId,
			MemberId:       item.MemberId,
			CouponCode:     item.CouponCode,
			MemberNickname: item.MemberNickname,
			GetType:        item.GetType,
			CreateTime:     item.CreateTime,
			UseStatus:      item.UseStatus,
			UseTime:        item.UseTime,
			OrderId:        item.OrderId,
			OrderSn:        item.OrderSn,
		})
	}

	return &types.ListCouponHistoryResp{
		Data:    list,
		Code:    0,
		Message: "获取会员优惠券历史列表成功",
	}, nil
}
