package main

import (
	infra_tools "templates-library/infra/tools"
)

func main() {
	logger := infra_tools.NewStdLogger()
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Success("This is a success message")
	logger.Error("This is an error message")
}
