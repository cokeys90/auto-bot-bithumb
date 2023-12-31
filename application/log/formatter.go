package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

const defaultTimeFormat = "2006-01-02T15:04:05.000Z07:00"

// PlainTextFormatter SpringBoot 같은 로깅 포맷팅
type PlainTextFormatter struct {
	logrus.Formatter
}

func (f *PlainTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := &bytes.Buffer{}
	level := strings.ToUpper(entry.Level.String())
	frame := entry.Caller

	functionPath := f.shortFunction(frame.Function)

	_, _ = fmt.Fprintf(b, "%-29s [%-5s] %s {%d} - %s", entry.Time.Format(defaultTimeFormat), level, functionPath, frame.Line, entry.Message)
	for k, v := range entry.Data {
		_, _ = fmt.Fprintf(b, " %s=%v", k, v)
	}
	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *PlainTextFormatter) shortFunction(_function string) (lastData string) {
	defer func() {
		if r := recover(); r != nil {
			lastData = _function
		}
	}()
	packages := strings.Split(_function, "/")
	newPackages := make([]string, len(packages))
	for index, val := range packages {
		if index == (len(packages) - 1) {
			break
		}
		newPackages[index] = string(rune(val[0]))
	}
	newPackages[len(newPackages)-1] = packages[len(packages)-1]
	lastData = strings.Join(newPackages, ".")
	return lastData
}
