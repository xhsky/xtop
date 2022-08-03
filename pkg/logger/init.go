package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	file, _ := os.OpenFile("xtop.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	writers := []io.Writer{
		file,
	}
	fileWriter := io.MultiWriter(writers...)
	log.SetOutput(fileWriter)
	log.SetLevel(log.InfoLevel)
}
