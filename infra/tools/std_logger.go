package tools

import (
	"fmt"
	"os"
	inter_infra_tools "templates-library/core/interfaces/infra/tools"
	infra_enums "templates-library/infra/enums"
)

type (
	StdLogger struct {
		log_level_config infra_enums.LogLevel
	}
)

func NewStdLogger() inter_infra_tools.ILogger {
	logLevelConfig := os.Getenv("BYT3UNION_LOG_LEVEL")
	allLevels := infra_enums.GetAllLogLevels()
	if logLevelConfig == "" {
		panic("BYT3UNION_LOG_LEVEL environment variable is not set. Accepted levels: " + fmt.Sprint(allLevels))
	}
	logLevelEnum, err := infra_enums.ConvertLogLevelFromStringToEnum(logLevelConfig)
	if err != nil {
		panic(err)
	}
	return &StdLogger{
		log_level_config: logLevelEnum,
	}
}

func (sl *StdLogger) Debug(msg string) {
	sl.writeLog(infra_enums.DEBUG, msg)
}

func (sl *StdLogger) Info(msg string) {
	sl.writeLog(infra_enums.INFO, msg)
}

func (sl *StdLogger) Warn(msg string) {
	sl.writeLog(infra_enums.WARN, msg)
}

func (sl *StdLogger) Success(msg string) {
	sl.writeLog(infra_enums.SUCCESS, msg)
}

func (sl *StdLogger) Error(msg string) {
	sl.writeLog(infra_enums.ERROR, msg)
}

func (sl *StdLogger) writeLog(level infra_enums.LogLevel, msg string) error {
	if level < sl.log_level_config {
		return nil
	}
	levelStr, err := infra_enums.ConvertLogLevelToString(level)
	if err != nil {
		return err
	}
	if level == infra_enums.ERROR {
		fmt.Fprintf(os.Stderr, "[%s]: %s\n", levelStr, msg)
		return nil
	}
	fmt.Fprintf(os.Stdout, "[%s]: %s\n", levelStr, msg)
	return nil
}
