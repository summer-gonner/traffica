package dictitemservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/summmer-gonner/traffica/sys/gen/model"
	"github.com/summmer-gonner/traffica/sys/gen/query"
	"github.com/summmer-gonner/traffica/sys/internal/svc"
	"github.com/summmer-gonner/traffica/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateDictItemLogic 更新字典数据
/*
Author: LiuFeiHua
Date: 2024/5/28 17:03
*/
type UpdateDictItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictItemLogic {
	return &UpdateDictItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateDictItem 更新字典数据
// 1.根据字典数据id查询字典数据是否已存在
// 2.字典数据存在时,则直接更新字典数据
func (l *UpdateDictItemLogic) UpdateDictItem(in *sysclient.UpdateDictItemReq) (*sysclient.UpdateDictItemResp, error) {
	q := query.SysDictItem.WithContext(l.ctx)

	// 1.根据字典数据id查询字典数据是否已存在
	_, err := q.Where(query.SysDictItem.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据字典数据id：%d,查询字典数据信息失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询字典数据信息失败"))
	}

	dictItem := &model.SysDictItem{
		ID:         in.Id,
		DictType:   in.DictType,
		DictLabel:  in.DictLabel,
		DictValue:  in.DictValue,
		DictStatus: in.DictStatus,
		DictSort:   in.DictSort,
		Remark:     in.Remark,
		IsDefault:  in.IsDefault,
		UpdateBy:   in.UpdateBy,
	}

	// 2.字典数据存在时,则直接更新字典数据
	_, err = q.Updates(dictItem)

	if err != nil {
		logc.Errorf(l.ctx, "更新字典数据信息失败,参数:%+v,异常:%s", dictItem, err.Error())
		return nil, errors.New(fmt.Sprintf("更新字典数据信息失败"))
	}

	return &sysclient.UpdateDictItemResp{}, nil
}
