package user

import (
	"context"
	"github.com/summer-gonner/traffica/admin/internal/common/errorx"
	"github.com/summer-gonner/traffica/admin/internal/svc"
	"github.com/summer-gonner/traffica/admin/internal/types"
	"github.com/summer-gonner/traffica/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryUserDetailLogic 查询用户详情
/*
Author: LiuFeiHua
Date: 2024/5/29 18:20
*/
type QueryUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserDetailLogic {
	return &QueryUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryUserDetail 查询用户详情
func (l *QueryUserDetailLogic) QueryUserDetail(req *types.QueryUserDetailReq) (resp *types.QueryUserDetailResp, err error) {
	item, err := l.svcCtx.UserService.QueryUserDetail(l.ctx, &sysclient.QueryUserDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询用户详情,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryUserDetailData{
		Avatar:     item.Avatar,
		CreateBy:   item.CreateBy,
		CreateTime: item.CreateTime,
		DeptId:     item.DeptId,
		Email:      item.Email,
		Id:         item.Id,
		LoginIp:    item.LoginIp,
		LoginTime:  item.LoginTime,
		Mobile:     item.Mobile,
		Nickname:   item.NickName,
		Remark:     item.Remark,
		UpdateBy:   item.UpdateBy,
		UpdateTime: item.UpdateTime,
		Username:   item.UserName,
		Status:     item.UserStatus,
		PostIds:    item.PostIds,
	}

	return &types.QueryUserDetailResp{
		Code:    "000000",
		Message: "查询用户详情成功",
		Data:    data,
	}, nil
}
