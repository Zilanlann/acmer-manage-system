package zap

import (
	"fmt"
	"os"

	"github.com/zilanlann/acmer-manage-system/server/pkg/zap/internal"
	"github.com/zilanlann/acmer-manage-system/server/setting"
	"github.com/zilanlann/acmer-manage-system/server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(setting.ZapSetting.Dir); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", setting.ZapSetting.Dir)
		_ = os.Mkdir(setting.ZapSetting.Dir, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if setting.ZapSetting.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
