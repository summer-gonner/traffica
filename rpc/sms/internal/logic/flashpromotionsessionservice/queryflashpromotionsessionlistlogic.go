package flashpromotionsessionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryFlashPromotionSessionListLogic 查询限时购场次表列表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:49
*/
type QueryFlashPromotionSessionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFlashPromotionSessionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFlashPromotionSessionListLogic {
	return &QueryFlashPromotionSessionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryFlashPromotionSessionList 查询限时购场次表列表
func (l *QueryFlashPromotionSessionListLogic) QueryFlashPromotionSessionList(in *smsclient.QueryFlashPromotionSessionListReq) (*smsclient.QueryFlashPromotionSessionListResp, error) {
	q := query.SmsFlashPromotionSession.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询限时购场次列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.FlashPromotionSessionListData
	for _, item := range result {

		list = append(list, &smsclient.FlashPromotionSessionListData{
			Id:         item.ID,
			Name:       item.Name,
			StartTime:  item.StartTime,
			EndTime:    item.EndTime,
			Status:     item.Status,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &smsclient.QueryFlashPromotionSessionListResp{
		Total: count,
		List:  list,
	}, nil

}
