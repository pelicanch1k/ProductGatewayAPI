package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

type writerHook struct {
	Writers   []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}

	for _, writer := range hook.Writers {
		writer.Write([]byte(line))
	}

	return nil
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func Init() {
	log := logrus.New()

	log.SetReportCaller(true)

	log.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	//err := os.MkdirAll("logs", 0644)
	//if err != nil {
	//	panic(err)
	//}
	//
	//allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	//if err != nil {
	//	panic(fmt.Sprintf("Failed to open logs file: %s", err))
	//}

	log.SetOutput(io.Discard)

	log.AddHook(&writerHook{
		Writers:   []io.Writer{os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	log.SetLevel(logrus.DebugLevel)
}
