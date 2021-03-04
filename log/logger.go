package log

import (
	"fmt"
	"gin_demo/config"
	"github.com/sirupsen/logrus"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"os"
	"path"
	"time"
)

var Log *logrus.Logger
var LogAccess *logrus.Logger

func LogInit() {
	logFilePath := ""
	logPath := config.Cfg.LogPath
	if len(logPath) == 0 {
		//获取当前目录
		if dir, err := os.Getwd(); err == nil {
			logFilePath = dir + "/logs/"
		}
	} else {
		//指定目录
		logFilePath = logPath + "/logs/"
	}

	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}

	rootLogInit(logFilePath)
	accessLogInit(logFilePath)
}

func rootLogInit(logFilePath string) {
	logFileName := "root.log"

	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	Log = logrus.New()
	//设置输出
	Log.Out = src
	Log.Out = os.Stdout
	//设置日志级别
	Log.SetLevel(logrus.DebugLevel)

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName + "-%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(2天)
		rotatelogs.WithMaxAge(2*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	//设置日志格式
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})

	// 新增 Hook
	Log.AddHook(lfHook)

}

func accessLogInit(logFilePath string) {
	logFileNameAccess := "access.log"

	fileNameAccess := path.Join(logFilePath, logFileNameAccess)
	if _, err := os.Stat(fileNameAccess); err != nil {
		if _, err := os.Create(fileNameAccess); err != nil {
			fmt.Println(err.Error())
		}
	}

	srcAccess, err := os.OpenFile(fileNameAccess, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	LogAccess = logrus.New()
	//设置输出
	LogAccess.Out = srcAccess
	LogAccess.Out = os.Stdout
	//设置日志级别
	LogAccess.SetLevel(logrus.DebugLevel)

	// 设置 rotatelogs
	logWriterAccess, err := rotatelogs.New(
		// 分割后的文件名称
		fileNameAccess + "-%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileNameAccess),

		// 设置最大保存时间(2天)
		rotatelogs.WithMaxAge(2*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMapAccess := lfshook.WriterMap{
		logrus.InfoLevel:  logWriterAccess,
		logrus.FatalLevel: logWriterAccess,
		logrus.DebugLevel: logWriterAccess,
		logrus.WarnLevel:  logWriterAccess,
		logrus.ErrorLevel: logWriterAccess,
		logrus.PanicLevel: logWriterAccess,
	}

	lfHookAccess := lfshook.NewHook(writeMapAccess, &logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})

	// 新增 Hook
	LogAccess.AddHook(lfHookAccess)
}
