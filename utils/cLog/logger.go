package cLog

import (
	"github.com/cokeys90/auto-bot-bithumb/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var logger *logrus.Logger

type Config struct {
	ZonedId  string       // 타임존
	Level    logrus.Level // 출력 로그 레벨
	FilePath string       // 파일경로 (파일명 제외)
}

func Setup(_config Config) {
	if logger == nil {
		logger = newLogger(_config)
	}
}

// Instance 신규 로거 생성
func Instance() *logrus.Logger {
	return logger
}

func DefaultConfig() Config {
	return Config{
		ZonedId:  "Asia/Seoul",
		Level:    logrus.InfoLevel,
		FilePath: "logs",
	}
}

// NewInstance 신규 로거 생성
func NewInstance(_config Config) *logrus.Logger {
	return newLogger(_config)
}

func newLogger(_config Config) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(new(utils.PlainTextFormatter))
	log.SetReportCaller(true)
	log.Level = _config.Level
	//cLog.AddHook(NewLoggerHook())

	isTzErr := false
	var tzErr error

	tz, tzErr := time.LoadLocation(_config.ZonedId)
	if tzErr != nil {
		tz = time.Local
		isTzErr = true
	}

	if _config.FilePath == "" {
		_config.FilePath = "logs"
	}
	filePath := _config.FilePath

	fileName := "/" + time.Now().In(tz).Format(time.DateOnly) + ".cLog"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		_ = os.Mkdir(filePath, 0755)
	}

	fullFilePath := _config.FilePath + fileName
	file, err := os.OpenFile(fullFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		writer := io.MultiWriter(os.Stdout, file)
		log.Out = writer
	} else {
		log.Out = os.Stdout
		log.Info("Failed to cLog to file, using default stdout")
	}

	if isTzErr {
		log.Warnf("타임존 설정오류: %s", tzErr.Error())
	}

	return log
}

// ToFields 구조체를 로깅 가능한 타입으로 변환
func ToFields(_struct any) logrus.Fields {
	fieldMap := make(logrus.Fields)
	_ = mapstructure.Decode(_struct, &fieldMap)
	return fieldMap
}

//// ConsoleLogHook 콘솔에 로깅을 하기 위한 훅
//type ConsoleLogHook struct {
//}
//
//func NewLoggerHook() *ConsoleLogHook {
//	hook := new(ConsoleLogHook)
//	return hook
//}
//
//func (h *ConsoleLogHook) Levels() []logrus.Level {
//	return []logrus.Level{
//		logrus.PanicLevel,
//		logrus.FatalLevel,
//		logrus.ErrorLevel,
//		logrus.WarnLevel,
//		logrus.InfoLevel,
//		logrus.DebugLevel,
//		logrus.TraceLevel,
//	}
//}
//
//func (h *ConsoleLogHook) Fire(entry *logrus.Entry) error {
//	str, _ := entry.String()
//	fmt.Print(str)
//	return nil
//}
