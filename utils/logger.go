package utils

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

type ILogger interface {
	Debug(msg string, params ...interface{}) string
	Info(msg string, params ...interface{}) string
	Success(msg string, params ...interface{}) string
	Warn(msg string, params ...interface{}) string
	Error(msg string, params ...interface{}) string
	Fatal(msg string, params ...interface{})
}

type logger struct {
	logStructureFormat string
	colorReset         string
	colorCyan          string
	colorRed           string
	colorYellow        string
	colorGreen         string
}

var GlobalLogger ILogger

func InitLogger() {
	GlobalLogger = NewLogger()
}

func NewLogger() ILogger {
	return &logger{
		logStructureFormat: "[%v] origin:[%s:%d] msg:[%v] ",
		colorReset:         "\033[0m",
		colorCyan:          "\033[96m",
		colorRed:           "\033[91m",
		colorYellow:        "\033[93m",
		colorGreen:         "\033[92m",
	}
}

func (l *logger) log(level, color, msg string, params ...interface{}) string {
	pc, _, line, _ := runtime.Caller(2)
	logResult := fmt.Sprintf(l.logStructureFormat, level, filepath.Base(runtime.FuncForPC(pc).Name()), line, msg)
	log.Printf(string(color)+logResult+string(l.colorReset), params...)
	return logResult
}

func (l *logger) Debug(msg string, params ...interface{}) string {
	return l.log("DEBU", string(l.colorReset), msg, params...)
}

func (l *logger) Info(msg string, params ...interface{}) string {
	return l.log("INFO", string(l.colorCyan), msg, params...)
}

func (l *logger) Success(msg string, params ...interface{}) string {
	formattedMsg := fmt.Sprintf(msg, params...)
	log.Printf(string(l.colorGreen)+"[SUCC] %s"+string(l.colorReset), formattedMsg)
	return formattedMsg
}

func (l *logger) Warn(msg string, params ...interface{}) string {
	return l.log("WARN", string(l.colorYellow), msg, params...)
}

func (l *logger) Error(msg string, params ...interface{}) string {
	return l.log("ERRO", string(l.colorRed), msg, params...)
}

func (l *logger) Fatal(msg string, params ...interface{}) {
	logResult := l.log("FATA", string(l.colorRed), msg, params...)
	log.Panic(logResult)
}
