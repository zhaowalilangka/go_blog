package global

import (
	"github.com/zhaowalilangka/go_blog/pkg/errcode/setting"
	"github.com/zhaowalilangka/go_blog/pkg/logger"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
