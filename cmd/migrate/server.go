package migrate

import (
	"github.com/cihub/seelog"
	"github.com/spf13/cobra"
	"tcod/database"
	"tcod/models/gorm"
	"tcod/tools"
	"tcod/tools/config"
)

var (
	configFile   string
	mode     string
	StartCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize the database",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "conf/settings.yml", "start tcod with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "boot mode ; eg:dev,test,prod")
}

func migrateModel() error {
	if config.DatabaseConfig.Dbtype == "mysql" {
		database.Eloquent = database.Eloquent.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	}
	return gorm.AutoMigrate(database.Eloquent)
}

func run() {
	//1. 设置日志
	tools.InitLog()

	//2. 读取配置
	config.ConfigSetup(configFile)
	defer database.Eloquent.Close()

	//3. 初始化数据库链接
	database.Setup()

	//4. 数据库迁移
	err := migrateModel()
	if err != nil {
		panic("数据库结构初始化失败: " + err.Error())
	}
	seelog.Info("数据库结构初始化成功! ")
}