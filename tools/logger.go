package tools

import (
	"github.com/cihub/seelog"
	"os"
)

func InitLog() {
	defer seelog.Flush()
	logPath,_ := os.Getwd()
	logger, err := seelog.LoggerFromConfigAsFile(logPath + "/conf/logger.xml")
	if err != nil {
		seelog.Critical("Err parsing log config file", err)
		os.Exit(0)
	}
	seelog.ReplaceLogger(logger)
}