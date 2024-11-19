package orderservicelogic

import (
	"context"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// OrderAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/13 9:29
*/
type OrderAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderAddLogic {
	return &OrderAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderAdd 创建订单
// 1.插入order表
// 2.插入order_item表
// 2.删除cart_item表(删除购物车中的下单商品)
func (l *OrderAddLogic) OrderAdd(in *omsclient.OrderAddReq) (*omsclient.OrderAddResp, error) {

	// 1.插入order表
	orderInfo := buildOrderInfo(in)
	logc.Infof(l.ctx, "插入order表,参数：%+v", orderInfo)
	err := query.OmsOrder.WithContext(l.ctx).Create(orderInfo)
	if err != nil {
		logc.Errorf(l.ctx, "插入order表失败,参数：%+v,异常：%s", orderInfo, err.Error())
		return nil, err
	}

	// 获取订单id
	orderId := orderInfo.ID
	// 2.插入order_item表
	for _, orderItem := range in.OrderItemList {
		buildOrderItem(l, orderInfo, orderItem)
	}

	// 3.删除cart_item表(删除购物车中的下单商品)
	logc.Infof(l.ctx, "删除购物车中的下单商品,参数ids：%+v", in.CartItemIds)
	item := query.OmsCartItem
	_, err = item.WithContext(l.ctx).Where(item.MemberID.Eq(in.MemberId), item.ID.In(in.CartItemIds...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除购物车中的下单商品失败,参数：%+v,异常：%s", in.CartItemIds, err.Error())
		return nil, err
	}

	return &omsclient.OrderAddResp{
		Id: orderId,
	}, nil
}

// 2.构建下单商品信息
func buildOrderItem(l *OrderAddLogic, orderInfo *model.OmsOrder, orderItem *omsclient.OrderItemData) {
	err := query.OmsOrderItem.WithContext(l.ctx).Create(&model.OmsOrderItem{
		OrderID:           orderInfo.ID,
		OrderSn:           orderInfo.OrderSn,
		ProductID:         orderItem.ProductId,
		ProductPic:        orderItem.ProductPic,
		ProductName:       orderItem.ProductName,
		ProductBrand:      orderItem.ProductBrand,
		ProductSn:         orderItem.ProductSn,
		ProductPrice:      orderItem.ProductPrice,
		ProductQuantity:   orderItem.ProductQuantity,
		ProductSkuID:      orderItem.ProductSkuId,
		ProductSkuCode:    orderItem.ProductSkuCode,
		ProductCategoryID: orderItem.ProductCategoryId,
		PromotionName:     orderItem.PromotionName,
		PromotionAmount:   orderItem.PromotionAmount,
		CouponAmount:      orderItem.CouponAmount,
		IntegrationAmount: orderItem.IntegrationAmount,
		RealAmount:        orderItem.RealAmount,
		GiftIntegration:   orderItem.GiftIntegration,
		GiftGrowth:        orderItem.GiftGrowth,
		ProductAttr:       orderItem.ProductAttr,
	})

	if err != nil {
		logc.Errorf(l.ctx, "插入order_item失败,参数：%+v,异常：%s", orderItem, err.Error())
	}
}

// 1.构建订单信息
func buildOrderInfo(in *omsclient.OrderAddReq) *model.OmsOrder {
	return &model.OmsOrder{
		MemberID:              in.MemberId,
		CouponID:              in.CouponId,
		OrderSn:               in.OrderSn,
		CreateTime:            time.Now(),
		MemberUsername:        in.MemberUsername,
		TotalAmount:           in.TotalAmount,
		PayAmount:             in.PayAmount,
		FreightAmount:         in.FreightAmount,
		PromotionAmount:       in.PromotionAmount,
		IntegrationAmount:     in.IntegrationAmount,
		CouponAmount:          in.CouponAmount,
		DiscountAmount:        in.DiscountAmount,
		PayType:               in.PayType,
		SourceType:            in.SourceType,
		Status:                in.Status,
		OrderType:             in.OrderType,
		DeliveryCompany:       in.DeliveryCompany,
		DeliverySn:            in.DeliverySn,
		AutoConfirmDay:        in.AutoConfirmDay,
		Integration:           in.Integration,
		Growth:                in.Growth,
		PromotionInfo:         in.PromotionInfo,
		BillType:              in.BillType,
		BillHeader:            in.BillHeader,
		BillContent:           in.BillContent,
		BillReceiverPhone:     in.BillReceiverPhone,
		BillReceiverEmail:     in.BillReceiverEmail,
		ReceiverName:          in.ReceiverName,
		ReceiverPhone:         in.ReceiverPhone,
		ReceiverPostCode:      in.ReceiverPostCode,
		ReceiverProvince:      in.ReceiverProvince,
		ReceiverCity:          in.ReceiverCity,
		ReceiverRegion:        in.ReceiverRegion,
		ReceiverDetailAddress: in.ReceiverDetailAddress,
		Note:                  in.Note,
		ConfirmStatus:         in.ConfirmStatus,
		DeleteStatus:          in.DeleteStatus,
		UseIntegration:        in.UseIntegration,
	}
}
