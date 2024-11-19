package dictitemservicelogic

import (
	"context"
	"errors"
	"github.com/summmer-gonner/traffica/sys/gen/query"
	"github.com/summmer-gonner/traffica/sys/internal/svc"
	"github.com/summmer-gonner/traffica/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteDictItemLogic 删除字典数据
/*
Author: LiuFeiHua
Date: 2024/5/28 17:03
*/
type DeleteDictItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictItemLogic {
	return &DeleteDictItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteDictItem 删除字典数据
func (l *DeleteDictItemLogic) DeleteDictItem(in *sysclient.DeleteDictItemReq) (*sysclient.DeleteDictItemResp, error) {
	q := query.SysDictItem

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除字典数据信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除字典数据信息失败")
	}

	return &sysclient.DeleteDictItemResp{}, nil
}
