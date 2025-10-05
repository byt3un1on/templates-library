package tools

type (
	ILogger interface {
		Debug(msg string)
		Info(msg string)
		Warn(msg string)
		Success(msg string)
		Error(msg string)
	}
)
