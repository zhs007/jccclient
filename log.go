package jccclient

// FuncLog - log
type FuncLog func(level string, str string)

var funcLog FuncLog = func(level string, str string) {}

func outputLog(level string, str string) {
	funcLog(level, str)
}

// SetFuncLog - set funcLog
func SetFuncLog(flog FuncLog) {
	funcLog = flog
}
