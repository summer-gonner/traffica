package order

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderListLogic 订单
/*
Author: LiuFeiHua
Date: 2024/5/14 16:37
*/
type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderListLogic {
	return OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// OrderList 查询订单
func (l *OrderListLogic) OrderList(req types.ListOrderReq) (*types.ListOrderResp, error) {
	resp, err := l.svcCtx.OrderService.OrderList(l.ctx, &omsclient.OrderListReq{
		Current:         req.Current,
		PageSize:        req.PageSize,
		OrderSn:         strings.TrimSpace(req.OrderSn),
		MemberUsername:  strings.TrimSpace(req.MemberUsername),
		PayType:         req.PayType,
		SourceType:      req.SourceType,
		Status:          req.Status,
		OrderType:       req.OrderType,
		ReceiverKeyword: strings.TrimSpace(req.ReceiverKeyword),
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询订单信息列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询订单信息失败")
	}

	var list []*types.ListOrderData

	for _, item := range resp.List {

		list = append(list, &types.ListOrderData{
			Id:                    item.Id,
			MemberId:              item.MemberId,
			CouponId:              item.CouponId,
			OrderSn:               item.OrderSn,
			CreateTime:            item.CreateTime,
			MemberUsername:        item.MemberUsername,
			TotalAmount:           item.TotalAmount,
			PayAmount:             item.PayAmount,
			FreightAmount:         item.FreightAmount,
			PromotionAmount:       item.PromotionAmount,
			IntegrationAmount:     item.IntegrationAmount,
			CouponAmount:          item.CouponAmount,
			DiscountAmount:        item.DiscountAmount,
			PayType:               item.PayType,
			SourceType:            item.SourceType,
			Status:                item.Status,
			OrderType:             item.OrderType,
			DeliveryCompany:       item.DeliveryCompany,
			DeliverySn:            item.DeliverySn,
			AutoConfirmDay:        item.AutoConfirmDay,
			Integration:           item.Integration,
			Growth:                item.Growth,
			PromotionInfo:         item.PromotionInfo,
			BillType:              item.BillType,
			BillHeader:            item.BillHeader,
			BillContent:           item.BillContent,
			BillReceiverPhone:     item.BillReceiverPhone,
			BillReceiverEmail:     item.BillReceiverEmail,
			ReceiverName:          item.ReceiverName,
			ReceiverPhone:         item.ReceiverPhone,
			ReceiverPostCode:      item.ReceiverPostCode,
			ReceiverProvince:      item.ReceiverProvince,
			ReceiverCity:          item.ReceiverCity,
			ReceiverRegion:        item.ReceiverRegion,
			ReceiverDetailAddress: item.ReceiverDetailAddress,
			Note:                  item.Note,
			ConfirmStatus:         item.ConfirmStatus,
			DeleteStatus:          item.DeleteStatus,
			UseIntegration:        item.UseIntegration,
			PaymentTime:           item.PaymentTime,
			DeliveryTime:          item.DeliveryTime,
			ReceiveTime:           item.ReceiveTime,
			CommentTime:           item.CommentTime,
			ModifyTime:            item.ModifyTime,
			ListOperateHistory:    queryHistoryRecord(item),
			ListOrderItem:         queryOrderItems(item),
		})
	}

	return &types.ListOrderResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询订单信息成功",
	}, nil
}

// 获取商品项
func queryOrderItems(item1 *omsclient.OrderListData) []types.ListOrderItemData {
	var itemListData []types.ListOrderItemData
	for _, item := range item1.ItemListData {
		itemListData = append(itemListData, types.ListOrderItemData{
			Id:                item.Id,
			OrderId:           item.OrderId,
			OrderSn:           item.OrderSn,
			ProductId:         item.ProductId,
			ProductPic:        item.ProductPic,
			ProductName:       item.ProductName,
			ProductBrand:      item.ProductBrand,
			ProductSn:         item.ProductSn,
			ProductPrice:      item.ProductPrice,
			ProductQuantity:   item.ProductQuantity,
			ProductSkuId:      item.ProductSkuId,
			ProductSkuCode:    item.ProductSkuCode,
			ProductCategoryId: item.ProductCategoryId,
			PromotionName:     item.PromotionName,
			PromotionAmount:   item.PromotionAmount,
			CouponAmount:      item.CouponAmount,
			IntegrationAmount: item.IntegrationAmount,
			RealAmount:        item.RealAmount,
			GiftIntegration:   item.GiftIntegration,
			GiftGrowth:        item.GiftGrowth,
			ProductAttr:       item.ProductAttr,
		})
	}

	return itemListData
}

// 获取操作记录
func queryHistoryRecord(item *omsclient.OrderListData) []types.OperateHistoryListData {
	var historyListData []types.OperateHistoryListData
	for _, operateHistory := range item.HistoryListData {
		historyListData = append(historyListData, types.OperateHistoryListData{
			Id:          operateHistory.Id,
			OrderId:     operateHistory.OrderId,
			OperateMan:  operateHistory.OperateMan,
			CreateTime:  operateHistory.CreateTime,
			OrderStatus: operateHistory.OrderStatus,
			Note:        operateHistory.Note,
		})
	}
	return historyListData
}
