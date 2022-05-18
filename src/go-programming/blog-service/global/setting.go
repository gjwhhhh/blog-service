package global

import (
	"github.com/go-programming/blog-service/pkg/logger"
	"github.com/go-programming/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
