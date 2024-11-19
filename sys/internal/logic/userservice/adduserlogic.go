package userservicelogic

import (
	"context"
	"errors"
	"fmt"
	model2 "github.com/summmer-gonner/traffica/sys/gen/model"
	"github.com/summmer-gonner/traffica/sys/gen/query"
	"github.com/summmer-gonner/traffica/sys/internal/svc"
	"github.com/summmer-gonner/traffica/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddUserLogic 新增用户
/*
Author: LiuFeiHua
Date: 2023/12/18 14:13
*/
type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddUser 新增用户
// 1.根据用户名称查询用户是否已存在
// 2.如果用户已存在,则直接返回
// 3.用户不存在时,则直接添加用户
func (l *AddUserLogic) AddUser(in *sysclient.AddUserReq) (*sysclient.AddUserResp, error) {
	q := query.SysUser
	// 1.根据用户名称查询用户是否已存在
	userName := in.UserName
	count, err := q.WithContext(l.ctx).Where(q.UserName.Eq(userName)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据用户名称：%s,查询用户信息失败,异常:%s", userName, err.Error())
		return nil, errors.New(fmt.Sprintf("查询用户信息失败"))
	}

	//2.如果用户已存在,则直接返回
	if count > 0 {
		logc.Errorf(l.ctx, "用户信息已存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("用户：%s,已存在", userName))
	}

	//3.用户不存在时,则直接添加用户
	user := &model2.SysUser{
		UserName:   in.UserName,
		NickName:   in.NickName,
		Avatar:     in.Avatar,
		Password:   "123456",
		Salt:       "123456", //todo 待完善
		Email:      in.Email,
		Mobile:     in.Mobile,
		UserStatus: in.UserStatus,
		DeptID:     in.DeptId,
		Remark:     in.Remark,
		CreateBy:   in.CreateBy,
	}

	err = query.Q.Transaction(func(tx *query.Query) error {
		err = tx.SysUser.WithContext(l.ctx).Create(user)

		if err != nil {
			logc.Errorf(l.ctx, "新增用户异常,参数:%+v,异常:%s", user, err.Error())
			return err
		}

		var userPosts []*model2.SysUserPost
		for _, postId := range in.PostIds {
			userPosts = append(userPosts, &model2.SysUserPost{
				UserID: user.ID,
				PostID: postId,
			})
		}

		postDo := tx.SysUserPost.WithContext(l.ctx)
		//清空脏数字
		_, err = postDo.Where(tx.SysUserPost.UserID.Eq(user.ID)).Delete()
		if err != nil {
			logc.Errorf(l.ctx, "删除用户与岗位关联异常,参数:%+v,异常:%s", user, err.Error())
			return err
		}

		err = postDo.Create(userPosts...)
		if err != nil {
			logc.Errorf(l.ctx, "新增用户与岗位关联异常,参数:%+v,异常:%s", user, err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		logc.Errorf(l.ctx, "新增用户异常,参数:%+v,异常:%s", user, err.Error())
		return nil, errors.New("新增用户异常")
	}
	return &sysclient.AddUserResp{}, nil
}
