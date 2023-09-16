package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
)

// 包全局变量
var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	Logger          *logger.Logger
)
