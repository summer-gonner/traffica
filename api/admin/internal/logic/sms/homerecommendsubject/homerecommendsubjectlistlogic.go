package homerecommendsubject

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeRecommendSubjectListLogic 人气推荐专题
/*
Author: LiuFeiHua
Date: 2024/5/14 9:43
*/
type HomeRecommendSubjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendSubjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendSubjectListLogic {
	return HomeRecommendSubjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// HomeRecommendSubjectList 查询人气推荐专题
func (l *HomeRecommendSubjectListLogic) HomeRecommendSubjectList(req types.ListHomeRecommendSubjectReq) (*types.ListHomeRecommendSubjectResp, error) {
	resp, err := l.svcCtx.HomeRecommendSubjectService.QueryHomeRecommendSubjectList(l.ctx, &smsclient.QueryHomeRecommendSubjectListReq{
		PageNum:         req.Current,
		PageSize:        req.PageSize,
		SubjectName:     strings.TrimSpace(req.SubjectName),
		RecommendStatus: req.RecommendStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询人气推荐专题列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询人气推荐专题失败")
	}

	var list []*types.ListHomeRecommendSubjectData

	for _, item := range resp.List {
		list = append(list, &types.ListHomeRecommendSubjectData{
			Id:              item.Id,
			SubjectId:       item.SubjectId,
			SubjectName:     item.SubjectName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	return &types.ListHomeRecommendSubjectResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询人气推荐专题成功",
	}, nil
}
