package userservicelogic

import (
	"context"
	"errors"
	"github.com/summer-gonner/traffica/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"github.com/summer-gonner/traffica/sys/internal/svc"
	"github.com/summer-gonner/traffica/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserProfileLogic {
	return &UserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取账号资料
func (l *UserProfileLogic) UserProfile(in *sysclient.ProfileReq) (*sysclient.ProfileResp, error) {
	// todo: add your logic here and delete this line
	//1.根据id查询用户信息
	q := query.SysUser
	profile, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.UserId)).First()

	// 2.判断用户是否存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "用户不存在,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("用户不存在")
	}

	if err != nil {
		logc.Errorf(l.ctx, "查询用户信息,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户信息异常")
	}
	return &sysclient.ProfileResp{
		Nickname: profile.NickName,
		Email:    profile.Email,
		Username: profile.UserName,
		Phone:    profile.Mobile,
		Remark:   profile.Remark,
		Avatar:   profile.Avatar,
	}, nil
}
