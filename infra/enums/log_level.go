package enums

import "fmt"

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	SUCCESS
	WARN
	ERROR
)

func GetAllLogLevels() []LogLevel {
	return []LogLevel{DEBUG, INFO, WARN, SUCCESS, ERROR}
}

func GetAllLogLevelsAsString() ([]string, error) {
	getAllLogLevels := GetAllLogLevels()
	var logLevelsAsString []string
	for _, logLevel := range getAllLogLevels {
		logLevelAsStr, err := ConvertLogLevelToString(logLevel)
		if err != nil {
			return []string(nil), err
		}
		logLevelsAsString = append(logLevelsAsString, logLevelAsStr)
	}
	return logLevelsAsString, nil
}

func ConvertLogLevelFromStringToEnum(logLevelStr string) (LogLevel, error) {
	switch logLevelStr {
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "WARN":
		return WARN, nil
	case "SUCCESS":
		return SUCCESS, nil
	case "ERROR":
		return ERROR, nil
	default:
		allLogLevels, err := GetAllLogLevelsAsString()
		if err != nil {
			return -1, err
		}
		return -1, fmt.Errorf("invalid log level: %s. Accepted levels: %s", logLevelStr, allLogLevels)
	}
}

func ConvertLogLevelToString(logLevel LogLevel) (string, error) {
	switch logLevel {
	case DEBUG:
		return "DEBUG", nil
	case INFO:
		return "INFO", nil
	case WARN:
		return "WARN", nil
	case SUCCESS:
		return "SUCCESS", nil
	case ERROR:
		return "ERROR", nil
	default:
		return "", fmt.Errorf("invalid log level: %d", logLevel)
	}
}
