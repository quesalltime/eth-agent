package logger

import (
	"os"

	"eth-agent/config"

	"github.com/sirupsen/logrus"
)

var (
	console *logrus.Logger
	// consoleEntry *logrus.Entry

	file *logrus.Logger
	// fileEntry *logrus.Entry

	// validate *validator.Validate
)

// Console -
func Console() *logrus.Logger {
	if console != nil {
		return console
	}

	console = logrus.New()
	customFormatter := new(logrus.TextFormatter)
	customFormatter.FullTimestamp = true
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	//customFormatter.DisableTimestamp = false
	//customFormatter.DisableColors = false

	console.Formatter = customFormatter
	console.SetLevel(config.ToLogLevel(config.SysConf.EthProxy.LogLevel))
	console.Out = os.Stdout

	// ConsoleEntry = Console.WithFields(logrus.Fields{
	// 	"BartConsole": "Ethan",
	// })
	return console
}

// File -
func File() *logrus.Logger {
	if file != nil {
		return file
	}

	file = logrus.New()
	customFormatter := new(logrus.TextFormatter)
	customFormatter.FullTimestamp = true
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	//customFormatter.DisableTimestamp = false
	//customFormatter.DisableColors = false

	file.Formatter = customFormatter
	file.SetLevel(logrus.DebugLevel)
	logFile, err := os.OpenFile(config.SysConf.EthProxy.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err == nil {
		file.Out = logFile
	}

	// FileEntry = File.WithFields(logrus.Fields{
	// 	"BartFile": "Ethan",
	// })
	return file
}
