package core

import (
	"bytes"
	"fmt"
	"gin-admin-api/config"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

const (
	red =31
	yellow =33
	blue = 36
	gray = 37
)

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry)([]byte,error){
	var levelColor int
	switch entry.Level{
	case logrus.DebugLevel,logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel,logrus.FatalLevel,logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil{
		b = entry.Buffer
	}else{
		b = &bytes.Buffer{} 
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller(){
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d",path.Base(entry.Caller.File),entry.Caller.Line)

		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s: %s\n", 
			timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	}else{
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m: %s\n", 
		timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(),nil
}

func InitLogger() *logrus.Logger{
	mLog := logrus.New()
	mLog.SetOutput(os.Stdout) // 設置輸出類型
	mLog.SetReportCaller(config.Config.Logger.ShowLine) // 開啟返回行數名和行號
	mLog.SetFormatter(&LogFormatter{}) // 設置自定義的Formatter
	level,err := logrus.ParseLevel(config.Config.Logger.Level)

	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level) // 設置最低的level
	InitDefaultLogger() 
	return mLog
}

func InitDefaultLogger(){
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(config.Config.Logger.ShowLine)
	logrus.SetFormatter(&LogFormatter{})
	level,err := logrus.ParseLevel(config.Config.Logger.Level)
	if err != nil{
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}