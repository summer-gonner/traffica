package svc

import (
	"github.com/smartwalle/alipay/v3"
	"github.com/summmer-gonner/traffica/api/front/internal/config"
	"github.com/summmer-gonner/traffica/rpc/cms/client/preferredareaproductrelationservice"
	"github.com/summmer-gonner/traffica/rpc/cms/client/preferredareaservice"
	"github.com/summmer-gonner/traffica/rpc/cms/client/subjectproductrelationservice"
	"github.com/summmer-gonner/traffica/rpc/cms/client/subjectservice"
	"github.com/summmer-gonner/traffica/rpc/oms/client/cartitemservice"
	"github.com/summmer-gonner/traffica/rpc/oms/client/companyaddressservice"
	"github.com/summmer-gonner/traffica/rpc/oms/client/orderitemservice"
	"github.com/summmer-gonner/traffica/rpc/oms/client/orderoperatehistoryservice"
	"github.com/summmer-gonner/traffica/rpc/oms/client/orderreturnapplyservice"
	"github.com/summmer-gonner/traffica/rpc/oms/client/orderreturnreasonservice"
	"github.com/summmer-gonner/traffica/rpc/oms/client/orderservice"
	"github.com/summmer-gonner/traffica/rpc/oms/client/ordersettingservice"
	"github.com/summmer-gonner/traffica/rpc/pms/client/brandservice"
	"github.com/summmer-gonner/traffica/rpc/pms/client/commentreplayservice"
	"github.com/summmer-gonner/traffica/rpc/pms/client/commentservice"
	"github.com/summmer-gonner/traffica/rpc/pms/client/feighttemplateservice"
	"github.com/summmer-gonner/traffica/rpc/pms/client/productattributecategoryservice"
	"github.com/summmer-gonner/traffica/rpc/pms/client/productattributeservice"
	"github.com/summmer-gonner/traffica/rpc/pms/client/productattributevalueservice"
	"github.com/summmer-gonner/traffica/rpc/pms/client/productcategoryattributerelationservice"
	"github.com/summmer-gonner/traffica/rpc/pms/client/productcategoryservice"
	"github.com/summmer-gonner/traffica/rpc/pms/client/productfullreductionservice"
	"github.com/summmer-gonner/traffica/rpc/pms/client/productladderservice"
	"github.com/summmer-gonner/traffica/rpc/pms/client/productoperatelogservice"
	"github.com/summmer-gonner/traffica/rpc/pms/client/productservice"
	"github.com/summmer-gonner/traffica/rpc/pms/client/productvertifyrecordservice"
	"github.com/summmer-gonner/traffica/rpc/pms/client/skustockservice"
	"github.com/summmer-gonner/traffica/rpc/sms/client/couponhistoryservice"
	"github.com/summmer-gonner/traffica/rpc/sms/client/couponservice"
	"github.com/summmer-gonner/traffica/rpc/sms/client/flashpromotionlogservice"
	"github.com/summmer-gonner/traffica/rpc/sms/client/flashpromotionproductrelationservice"
	"github.com/summmer-gonner/traffica/rpc/sms/client/flashpromotionservice"
	"github.com/summmer-gonner/traffica/rpc/sms/client/flashpromotionsessionservice"
	"github.com/summmer-gonner/traffica/rpc/sms/client/homeadvertiseservice"
	"github.com/summmer-gonner/traffica/rpc/sms/client/homebrandservice"
	"github.com/summmer-gonner/traffica/rpc/sms/client/homenewproductservice"
	"github.com/summmer-gonner/traffica/rpc/sms/client/homerecommendproductservice"
	"github.com/summmer-gonner/traffica/rpc/sms/client/homerecommendsubjectservice"
	"github.com/summmer-gonner/traffica/rpc/ums/client/growthchangehistoryservice"
	"github.com/summmer-gonner/traffica/rpc/ums/client/integrationchangehistoryservice"
	"github.com/summmer-gonner/traffica/rpc/ums/client/integrationconsumesettingservice"
	"github.com/summmer-gonner/traffica/rpc/ums/client/memberbrandattentionservice"
	"github.com/summmer-gonner/traffica/rpc/ums/client/memberlevelservice"
	"github.com/summmer-gonner/traffica/rpc/ums/client/memberloginlogservice"
	"github.com/summmer-gonner/traffica/rpc/ums/client/membermembertagrelationservice"
	"github.com/summmer-gonner/traffica/rpc/ums/client/memberproductcategoryrelationservice"
	"github.com/summmer-gonner/traffica/rpc/ums/client/memberproductcollectionservice"
	"github.com/summmer-gonner/traffica/rpc/ums/client/memberreadhistoryservice"
	"github.com/summmer-gonner/traffica/rpc/ums/client/memberreceiveaddressservice"
	"github.com/summmer-gonner/traffica/rpc/ums/client/memberrulesettingservice"
	"github.com/summmer-gonner/traffica/rpc/ums/client/memberservice"
	"github.com/summmer-gonner/traffica/rpc/ums/client/memberstatisticsinfoservice"
	"github.com/summmer-gonner/traffica/rpc/ums/client/membertagservice"
	"github.com/summmer-gonner/traffica/rpc/ums/client/membertaskservice"
	"github.com/summmer-gonner/traffica/sys/client/deptservice"
	"github.com/summmer-gonner/traffica/sys/client/loginlogservice"
	"github.com/summmer-gonner/traffica/sys/client/menuservice"
	"github.com/summmer-gonner/traffica/sys/client/roleservice"
	"github.com/summmer-gonner/traffica/sys/client/userservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	//会员相关
	GrowthChangeHistoryService           growthchangehistoryservice.GrowthChangeHistoryService
	IntegrationChangeHistoryService      integrationchangehistoryservice.IntegrationChangeHistoryService
	IntegrationConsumeSettingService     integrationconsumesettingservice.IntegrationConsumeSettingService
	MemberLevelService                   memberlevelservice.MemberLevelService
	MemberLoginLogService                memberloginlogservice.MemberLoginLogService
	MemberMemberTagRelationService       membermembertagrelationservice.MemberMemberTagRelationService
	MemberProductCategoryRelationService memberproductcategoryrelationservice.MemberProductCategoryRelationService
	MemberProductCollectionService       memberproductcollectionservice.MemberProductCollectionService
	MemberReadHistoryService             memberreadhistoryservice.MemberReadHistoryService
	MemberReceiveAddressService          memberreceiveaddressservice.MemberReceiveAddressService
	MemberRuleSettingService             memberrulesettingservice.MemberRuleSettingService
	MemberService                        memberservice.MemberService
	MemberStatisticsInfoService          memberstatisticsinfoservice.MemberStatisticsInfoService
	MemberTagService                     membertagservice.MemberTagService
	MemberTaskService                    membertaskservice.MemberTaskService
	MemberBrandAttentionService          memberbrandattentionservice.MemberBrandAttentionService

	//系统相关
	DeptService     deptservice.DeptService
	LoginLogService loginlogservice.LoginLogService
	MenuService     menuservice.MenuService
	RoleService     roleservice.RoleService
	UserService     userservice.UserService
	//商品相关
	BrandService                            brandservice.BrandService
	CommentReplayService                    commentreplayservice.CommentReplayService
	CommentService                          commentservice.CommentService
	FeightTemplateService                   feighttemplateservice.FeightTemplateService
	ProductAttributeCategoryService         productattributecategoryservice.ProductAttributeCategoryService
	ProductAttributeService                 productattributeservice.ProductAttributeService
	ProductAttributeValueService            productattributevalueservice.ProductAttributeValueService
	ProductCategoryAttributeRelationService productcategoryattributerelationservice.ProductCategoryAttributeRelationService
	ProductCategoryService                  productcategoryservice.ProductCategoryService
	ProductFullReductionService             productfullreductionservice.ProductFullReductionService
	ProductLadderService                    productladderservice.ProductLadderService
	ProductOperateLogService                productoperatelogservice.ProductOperateLogService
	ProductService                          productservice.ProductService
	ProductVertifyRecordService             productvertifyrecordservice.ProductVertifyRecordService
	SkuStockService                         skustockservice.SkuStockService
	//订单相关
	CartItemService            cartitemservice.CartItemService
	CompanyAddressService      companyaddressservice.CompanyAddressService
	OrderItemService           orderitemservice.OrderItemService
	OrderOperateHistoryService orderoperatehistoryservice.OrderOperateHistoryService
	OrderReturnApplyService    orderreturnapplyservice.OrderReturnApplyService
	OrderReturnReasonService   orderreturnreasonservice.OrderReturnReasonService
	OrderService               orderservice.OrderService
	OrderSettingService        ordersettingservice.OrderSettingService
	//营销相关
	CouponHistoryService                 couponhistoryservice.CouponHistoryService
	CouponService                        couponservice.CouponService
	FlashPromotionLogService             flashpromotionlogservice.FlashPromotionLogService
	FlashPromotionProductRelationService flashpromotionproductrelationservice.FlashPromotionProductRelationService
	FlashPromotionService                flashpromotionservice.FlashPromotionService
	FlashPromotionSessionService         flashpromotionsessionservice.FlashPromotionSessionService
	HomeAdvertiseService                 homeadvertiseservice.HomeAdvertiseService
	HomeBrandService                     homebrandservice.HomeBrandService
	HomeNewProductService                homenewproductservice.HomeNewProductService
	HomeRecommendProductService          homerecommendproductservice.HomeRecommendProductService
	HomeRecommendSubjectService          homerecommendsubjectservice.HomeRecommendSubjectService
	//内容相关
	SubjectService                      subjectservice.SubjectService
	SubjectProductRelationService       subjectproductrelationservice.SubjectProductRelationService
	PreferredAreaService                preferredareaservice.PreferredAreaService
	PreferredAreaProductRelationService preferredareaproductrelationservice.PreferredAreaProductRelationService

	AlipayClient *alipay.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	//初始化支付宝客户端
	client, err := alipay.New(c.Alipay.AppId, c.Alipay.PrivateKey, c.Alipay.IsProduction)
	if err != nil {
		print("初始化支付宝失败")
	}
	umsClient := zrpc.MustNewClient(c.UmsRpc)
	sysClient := zrpc.MustNewClient(c.SysRpc)
	pmsClient := zrpc.MustNewClient(c.PmsRpc)
	omsClient := zrpc.MustNewClient(c.OmsRpc)
	smsClient := zrpc.MustNewClient(c.SmsRpc)
	cmsClient := zrpc.MustNewClient(c.CmsRpc)
	return &ServiceContext{
		Config:                               c,
		GrowthChangeHistoryService:           growthchangehistoryservice.NewGrowthChangeHistoryService(umsClient),
		IntegrationChangeHistoryService:      integrationchangehistoryservice.NewIntegrationChangeHistoryService(umsClient),
		IntegrationConsumeSettingService:     integrationconsumesettingservice.NewIntegrationConsumeSettingService(umsClient),
		MemberLevelService:                   memberlevelservice.NewMemberLevelService(umsClient),
		MemberLoginLogService:                memberloginlogservice.NewMemberLoginLogService(umsClient),
		MemberMemberTagRelationService:       membermembertagrelationservice.NewMemberMemberTagRelationService(umsClient),
		MemberProductCategoryRelationService: memberproductcategoryrelationservice.NewMemberProductCategoryRelationService(umsClient),
		MemberProductCollectionService:       memberproductcollectionservice.NewMemberProductCollectionService(umsClient),
		MemberReadHistoryService:             memberreadhistoryservice.NewMemberReadHistoryService(umsClient),
		MemberReceiveAddressService:          memberreceiveaddressservice.NewMemberReceiveAddressService(umsClient),
		MemberRuleSettingService:             memberrulesettingservice.NewMemberRuleSettingService(umsClient),
		MemberService:                        memberservice.NewMemberService(umsClient),
		MemberStatisticsInfoService:          memberstatisticsinfoservice.NewMemberStatisticsInfoService(umsClient),
		MemberTagService:                     membertagservice.NewMemberTagService(umsClient),
		MemberTaskService:                    membertaskservice.NewMemberTaskService(umsClient),
		MemberBrandAttentionService:          memberbrandattentionservice.NewMemberBrandAttentionService(umsClient),

		DeptService:     deptservice.NewDeptService(sysClient),
		LoginLogService: loginlogservice.NewLoginLogService(sysClient),
		MenuService:     menuservice.NewMenuService(sysClient),
		RoleService:     roleservice.NewRoleService(sysClient),
		UserService:     userservice.NewUserService(sysClient),

		BrandService:                            brandservice.NewBrandService(pmsClient),
		CommentReplayService:                    commentreplayservice.NewCommentReplayService(pmsClient),
		CommentService:                          commentservice.NewCommentService(pmsClient),
		FeightTemplateService:                   feighttemplateservice.NewFeightTemplateService(pmsClient),
		ProductAttributeCategoryService:         productattributecategoryservice.NewProductAttributeCategoryService(pmsClient),
		ProductAttributeService:                 productattributeservice.NewProductAttributeService(pmsClient),
		ProductAttributeValueService:            productattributevalueservice.NewProductAttributeValueService(pmsClient),
		ProductCategoryAttributeRelationService: productcategoryattributerelationservice.NewProductCategoryAttributeRelationService(pmsClient),
		ProductCategoryService:                  productcategoryservice.NewProductCategoryService(pmsClient),
		ProductFullReductionService:             productfullreductionservice.NewProductFullReductionService(pmsClient),
		ProductLadderService:                    productladderservice.NewProductLadderService(pmsClient),
		ProductOperateLogService:                productoperatelogservice.NewProductOperateLogService(pmsClient),
		ProductService:                          productservice.NewProductService(pmsClient),
		ProductVertifyRecordService:             productvertifyrecordservice.NewProductVertifyRecordService(pmsClient),
		SkuStockService:                         skustockservice.NewSkuStockService(pmsClient),

		CartItemService:            cartitemservice.NewCartItemService(omsClient),
		CompanyAddressService:      companyaddressservice.NewCompanyAddressService(omsClient),
		OrderItemService:           orderitemservice.NewOrderItemService(omsClient),
		OrderOperateHistoryService: orderoperatehistoryservice.NewOrderOperateHistoryService(omsClient),
		OrderReturnApplyService:    orderreturnapplyservice.NewOrderReturnApplyService(omsClient),
		OrderReturnReasonService:   orderreturnreasonservice.NewOrderReturnReasonService(omsClient),
		OrderService:               orderservice.NewOrderService(omsClient),
		OrderSettingService:        ordersettingservice.NewOrderSettingService(omsClient),

		CouponHistoryService:                 couponhistoryservice.NewCouponHistoryService(smsClient),
		CouponService:                        couponservice.NewCouponService(smsClient),
		FlashPromotionLogService:             flashpromotionlogservice.NewFlashPromotionLogService(smsClient),
		FlashPromotionProductRelationService: flashpromotionproductrelationservice.NewFlashPromotionProductRelationService(smsClient),
		FlashPromotionService:                flashpromotionservice.NewFlashPromotionService(smsClient),
		FlashPromotionSessionService:         flashpromotionsessionservice.NewFlashPromotionSessionService(smsClient),
		HomeAdvertiseService:                 homeadvertiseservice.NewHomeAdvertiseService(smsClient),
		HomeBrandService:                     homebrandservice.NewHomeBrandService(smsClient),
		HomeNewProductService:                homenewproductservice.NewHomeNewProductService(smsClient),
		HomeRecommendProductService:          homerecommendproductservice.NewHomeRecommendProductService(smsClient),
		HomeRecommendSubjectService:          homerecommendsubjectservice.NewHomeRecommendSubjectService(smsClient),

		SubjectService:                      subjectservice.NewSubjectService(cmsClient),
		SubjectProductRelationService:       subjectproductrelationservice.NewSubjectProductRelationService(cmsClient),
		PreferredAreaService:                preferredareaservice.NewPreferredAreaService(cmsClient),
		PreferredAreaProductRelationService: preferredareaproductrelationservice.NewPreferredAreaProductRelationService(cmsClient),

		AlipayClient: client,
	}
}
