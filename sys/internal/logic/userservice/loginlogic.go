package userservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	model2 "github.com/summer-gonner/traffica/sys/gen/model"
	"github.com/summer-gonner/traffica/sys/gen/query"
	"github.com/summer-gonner/traffica/sys/internal/logic/common"
	"github.com/summer-gonner/traffica/sys/internal/svc"
	"github.com/summer-gonner/traffica/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// LoginLogic 用户登录
/*
Author: LiuFeiHua
Date: 2023/12/18 14:08
*/
type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Login 根据用户名和密码登录
// 1.判断用户是否存在
// 2.判断密码是否正确
// 3.查询权限
// 4.获取部门信息
// 5.生成token
// 6.保存登录日志
// 7.更新登录时间
func (l *LoginLogic) Login(in *sysclient.LoginReq) (*sysclient.LoginResp, error) {
	q := query.SysUser
	user, err := q.WithContext(l.ctx).Where(q.UserName.Eq(in.Account)).Or(q.Mobile.Eq(in.Account)).First()

	// 1.判断用户是否存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		l.savaLoginLog(in, "error", fmt.Sprintf("用户不存在: %+v", in))
		logc.Errorf(l.ctx, "用户不存在,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("用户不存在")
	}

	if err != nil {
		l.savaLoginLog(in, "error", fmt.Sprintf("查询用户信息异常: %+v", in))
		logc.Errorf(l.ctx, "查询用户信息,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户信息异常")
	}

	// 2.判断密码是否正确
	if user.Password != in.Password {
		l.savaLoginLog(in, "error", fmt.Sprintf("用户密码不正确: %+v", in))
		logc.Errorf(l.ctx, "用户密码不正确,参数:%s", in.Password)
		return nil, errors.New("用户密码不正确")
	}

	// 3.查询权限
	apiUrls := l.queryApiUrls(user)

	if len(apiUrls) == 0 {
		l.savaLoginLog(in, "error", fmt.Sprintf("用户还没有分配角色或者还没有分配角色权限: %+v", in))
		logc.Errorf(l.ctx, "用户还没有分配角色或者还没有分配角色权限,参数:%+v", in)
		return nil, errors.New("用户还没有分配角色或者还没有分配角色权限")
	}

	// 4.获取部门信息
	sysDept := query.SysDept
	dept, err := sysDept.WithContext(l.ctx).Select(sysDept.DeptName).Where(sysDept.ID.Eq(user.DeptID)).First()
	if err != nil {
		l.savaLoginLog(in, "error", fmt.Sprintf("查询用户部门信息异常: %+v", in))
		logc.Errorf(l.ctx, "查询用户部门信息异常,参数deptId:%d", user.DeptID)
		return nil, errors.New("查询用户部门信息异常")
	}
	// 5.生成token
	jwtToken, err := l.getJwtToken(user.ID, user.DeptID, user.UserName, dept.DeptName)

	if err != nil {
		l.savaLoginLog(in, "error", fmt.Sprintf("生成token失败: %+v", in))
		logc.Errorf(l.ctx, "生成token失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("生成token失败")
	}

	// 6.保存登录日志
	l.savaLoginLog(in, "success", "登录成功")

	// 7.更新登录时间
	now := time.Now()
	_, _ = q.WithContext(l.ctx).Where(q.ID.Eq(user.ID)).Updates(&model2.SysUser{
		LoginTime:    &now,
		LoginIP:      in.IpAddress,
		LoginOs:      in.Os,
		LoginBrowser: in.Browser,
	})

	return &sysclient.LoginResp{
		Id:          user.ID,
		UserName:    user.UserName,
		AccessToken: jwtToken,
		ApiUrls:     apiUrls,
	}, nil
}

// 3.查询权限(判断是不是超级管理员：角色is_admin:1是超级管理员的角色)
func (l *LoginLogic) queryApiUrls(user *model2.SysUser) []string {
	db := l.svcCtx.DB

	var apiUrls []string
	// 4.1判断是不是超级管理员，则是超级管理员，拿所有权限
	if common.IsAdmin(l.ctx, user.ID, l.svcCtx.DB) {
		sql := `select background_url
			from sys_menu
			where background_url != ''`
		db.WithContext(l.ctx).Raw(sql).Select("background_url").Scan(&apiUrls)
	} else {
		sql := `select sm.background_url
			from sys_user_role sur
					 left join sys_role sr on sur.role_id = sr.id
					 left join sys_role_menu srm on sr.id = srm.role_id
					 left join sys_menu sm on srm.menu_id = sm.id
			where sur.user_id = ?`
		db.WithContext(l.ctx).Raw(sql, user.ID).Select("background_url").Scan(&apiUrls)
	}
	return apiUrls
}

// 保存登录日志
func (l *LoginLogic) savaLoginLog(in *sysclient.LoginReq, loginStatus, errorMsg string) {
	_ = query.SysLoginLog.WithContext(l.ctx).Create(&model2.SysLoginLog{
		UserName:    in.Account,
		LoginStatus: loginStatus,
		IPAddress:   in.IpAddress,
		Browser:     in.Browser,
		Os:          in.Os,
		ErrorMsg:    errorMsg,
		LoginTime:   time.Now(),
	})
}

// 生成jwt的token
func (l *LoginLogic) getJwtToken(userId, deptID int64, userName, deptName string) (string, error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	accessSecret := l.svcCtx.Config.JWT.AccessSecret

	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire
	claims["iat"] = now
	claims["userId"] = userId
	claims["userName"] = userName
	claims["deptID"] = deptID
	claims["deptName"] = deptName
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(accessSecret))
}
