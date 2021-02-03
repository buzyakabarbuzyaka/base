package logger

import (
	"github.com/buzyakabarbuzyaka/base/kit/config"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
	"os"
)

func Init(c config.LoggerConfig) *logrus.Logger {
	log := logrus.New()
	log.SetLevel(c.GetLevel())
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{})

	logFile, err := os.OpenFile(c.AllOutPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		log.Panic(err)
	}

	errFile, err := os.OpenFile(c.ErrorOutPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		log.Panic(err)
	}

	log.AddHook(&writer.Hook{ // Send logs with level higher than warning to allfile
		Writer: logFile,
		LogLevels: []logrus.Level{
			logrus.DebugLevel,
		},
	})

	log.AddHook(&writer.Hook{ // Send logs with level higher than warning to errFile
		Writer: errFile,
		LogLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		},
	})

	log.SetOutput(os.Stdout)
	return log
}
