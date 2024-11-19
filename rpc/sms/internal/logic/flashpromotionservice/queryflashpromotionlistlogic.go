package flashpromotionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryFlashPromotionListLogic 查询限时购表列表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:43
*/
type QueryFlashPromotionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFlashPromotionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFlashPromotionListLogic {
	return &QueryFlashPromotionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryFlashPromotionList 查询限时购表列表
func (l *QueryFlashPromotionListLogic) QueryFlashPromotionList(in *smsclient.QueryFlashPromotionListReq) (*smsclient.QueryFlashPromotionListResp, error) {
	q := query.SmsFlashPromotion.WithContext(l.ctx)
	if len(in.Title) > 0 {
		q = q.Where(query.SmsFlashPromotion.Title.Like("%" + in.Title + "%"))
	}

	if in.Status != 2 {
		q = q.Where(query.SmsFlashPromotion.Status.Eq(in.Status))
	}

	if len(in.StartDate) > 0 {
		startDate, _ := time.Parse("2006-01-02", in.StartDate)
		q = q.Where(query.SmsFlashPromotion.StartDate.Gte(startDate))
	}
	if len(in.EndDate) > 0 {
		endDate, _ := time.Parse("2006-01-02", in.EndDate)
		q = q.Where(query.SmsFlashPromotion.EndDate.Lte(endDate))
	}
	//if len(in.StartDate) > 0 {
	//		where = where + fmt.Sprintf(" AND start_date >= '%s'", in.StartDate)
	//	}
	//	if len(in.EndDate) > 0 {
	//		where = where + fmt.Sprintf(" AND end_date <= '%s'", in.EndDate)
	//	}
	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询限时购列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.FlashPromotionListData
	for _, item := range result {

		list = append(list, &smsclient.FlashPromotionListData{
			Id:         item.ID,
			Title:      item.Title,
			StartDate:  item.StartDate.Format("2006-01-02"),
			EndDate:    item.EndDate.Format("2006-01-02"),
			Status:     item.Status,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &smsclient.QueryFlashPromotionListResp{
		Total: count,
		List:  list,
	}, nil

}
