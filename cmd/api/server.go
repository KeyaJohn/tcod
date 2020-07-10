package api

import (
	"github.com/cihub/seelog"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"tcod/database"
	"tcod/tools"
	"tcod/tools/config"
)

var (
	configFile   string
	port     string
	mode     string
	StartCmd = &cobra.Command{
		Use:     "start",
		Short:   "Boot tcod",
		Example: "tcod client conf/settings.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "conf/settings.yml", "start tcod with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "boot mode ; eg:dev,test,prod")
}

func usage() {
	usageStr := `Boot tcod`
	seelog.Info("%s\n", usageStr)
}

func setup() {
	//1. 设置日志
	tools.InitLog()

	//2. 读取配置
	config.ConfigSetup(configFile)

	//3. 初始化数据库链接
	database.Setup()

}

func run() error {
	defer database.Eloquent.Close()

	//优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt,os.Kill)
	<-quit
	seelog.Info("%s Shutdown Collect.... Time:", tools.GetCurrntTimeStr())
	return nil
}
