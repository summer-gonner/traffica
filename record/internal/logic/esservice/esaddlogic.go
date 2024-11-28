package esservicelogic

import (
	"context"
	"errors"
	"fmt"
	model2 "github.com/summer-gonner/traffica/record/gen/model"
	"github.com/summer-gonner/traffica/record/gen/query"
	"github.com/summer-gonner/traffica/record/internal/svc"
	"github.com/summer-gonner/traffica/record/recordclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// EsAddLogic 新增es
type EsAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEsAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EsAddLogic {
	return &EsAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EsAddLogic) EsAdd(in *recordclient.EsAddReq) (*recordclient.EsAddResp, error) {

	q := query.RecEsInfo
	// 1.根据地址查询用户是否已存在
	address := in.Address
	count, err := q.WithContext(l.ctx).Where(q.Address.Eq(address)).Count()
	if err != nil {
		logc.Errorf(l.ctx, "根据es地址：%s,查询es表失败,异常：%s", address, err.Error())

		return nil, errors.New(fmt.Sprintf("查询es信息表失败"))
	}
	//2.如果es已经存在，则直接返回
	if count > 0 {
		logc.Errorf(l.ctx, "用户信息已存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("es:%s,已存在", address))
	}
	//3.es不存在，则直接添加es
	esInfo := &model2.RecEsInfo{
		Address:  in.Address,
		Username: in.Username,
		Password: in.Password,
		Name:     in.Name,
	}
	err = query.Q.Transaction(func(tx *query.Query) error {
		err := tx.RecEsInfo.WithContext(l.ctx).Create(esInfo)
		if err != nil {
			logc.Errorf(l.ctx, "新增es异常,参数:%+v,异常:%s", esInfo, err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &recordclient.EsAddResp{
		Result:  true,
		Message: "新增成功",
	}, nil
}
