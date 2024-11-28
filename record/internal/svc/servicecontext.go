package svc

import (
	"github.com/summer-gonner/traffica/record/gen/query"
	"github.com/summer-gonner/traffica/record/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	DB, err := gorm.Open(mysql.Open(c.Mysql.Datasource), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 settingLogConfig(),
	})
	if err != nil {
		panic(err)
	}

	logx.Debug("mysql已连接")
	query.SetDefault(DB)

	//sqlConn := sqlx.NewMysql(c.Mysql.Datasource)

	return &ServiceContext{
		Config: c,
		DB:     DB,
		//UserModel:     sysmodel.NewSysUserModel(sqlConn),
		//UserRoleModel: sysmodel.NewSysUserRoleModel(sqlConn),
		//RoleModel:     sysmodel.NewSysRoleModel(sqlConn),
		//RoleMenuModel: sysmodel.NewSysRoleMenuModel(sqlConn),
		//MenuModel:     sysmodel.NewSysMenuModel(sqlConn),
		//DictModel:     sysmodel.NewSysDictModel(sqlConn),
		//DeptModel:     sysmodel.NewSysDeptModel(sqlConn),
		//LoginLogModel: sysmodel.NewSysLoginLogModel(sqlConn),
		//SysLogModel:   sysmodel.NewSysLogModel(sqlConn),
		//ConfigModel:   sysmodel.NewSysConfigModel(sqlConn),
		//JobModel:      sysmodel.NewSysJobModel(sqlConn),
	}
}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	logx.Infof(format, args...)
}

// init log config
func settingLogConfig() logger.Interface {
	newLogger := logger.New(
		Writer{},
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)
	return newLogger
}
