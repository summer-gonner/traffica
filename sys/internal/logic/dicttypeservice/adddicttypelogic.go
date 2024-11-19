package dicttypeservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/summmer-gonner/traffica/sys/gen/model"
	"github.com/summmer-gonner/traffica/sys/gen/query"
	"github.com/summmer-gonner/traffica/sys/internal/svc"
	"github.com/summmer-gonner/traffica/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddDictTypeLogic 添加字典信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:02
*/
type AddDictTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDictTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDictTypeLogic {
	return &AddDictTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddDictType 添加字典信息
// 1.根据字典名称或者类型查询字典是否已存在
// 2.如果字典已存在,则直接返回
// 3.字典不存在时,则直接添加字典
func (l *AddDictTypeLogic) AddDictType(in *sysclient.AddDictTypeReq) (*sysclient.AddDictTypeResp, error) {
	q := query.SysDictType

	// 1.根据字典名称或者类型查询字典是否已存在
	count, err := q.WithContext(l.ctx).Where(q.DictName.Eq(in.DictName)).Or(q.DictType.Eq(in.DictType)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据字典名称：%+v,查询字典信息失败,异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("查询字典信息失败"))
	}

	//2.如果字典已存在,则直接返回
	if count > 0 {
		logc.Errorf(l.ctx, "字典信息已存在：%+v", in)
		return nil, errors.New("字典已存在")
	}

	dict := &model.SysDictType{
		DictName:   in.DictName,
		DictType:   in.DictType,
		DictStatus: in.DictStatus,
		Remark:     in.Remark,
		IsSystem:   in.IsSystem,
		IsDeleted:  0,
		CreateBy:   in.CreateBy,
		CreateTime: time.Now(),
	}
	err = q.WithContext(l.ctx).Create(dict)

	if err != nil {
		logc.Errorf(l.ctx, "添加字典信息失败,参数:%+v,异常:%s", dict, err.Error())
		return nil, errors.New("添加字典信息失败")
	}

	return &sysclient.AddDictTypeResp{}, nil
}
