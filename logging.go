package logging

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

//goland:noinspection GoUnnecessarilyExportedIdentifiers,GoUnusedConst
const (
	Command     = "[COMMAND  ]"
	Cron        = "[CRON     ]"
	Database    = "[DATABASE ]"
	Discord     = "[DISCORD  ]"
	Picarto     = "[PICARTO  ]"
	Trovo       = "[TROVO    ]"
	Twitch      = "[TWITCH   ]"
	YouTubeLive = "[YOUTUBE-L]"
	YouTubePub  = "[YOUTUBE-P]"
)

var (
	log = logrus.Logger{
		Out: os.Stderr,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %msg%\n",
		},
		ReportCaller: true,
	}
)

func init() {
	//goland:noinspection GoBoolExpressions
	if runtime.GOOS == "windows" {
		log.SetLevel(logrus.DebugLevel)

	} else {
		log.SetLevel(logrus.InfoLevel)
	}
}

// Traceln - Logs the event at the Trace level
//
//goland:noinspection GoUnusedExportedFunction
func Traceln(args ...any) {
	if log.Level == logrus.TraceLevel {
		log.Traceln(args...)
	}
}

// Tracef - Logs the event at the Traceln level with formatting
//
//goland:noinspection GoUnusedExportedFunction
func Tracef(format string, args ...any) {
	if log.Level == logrus.TraceLevel {
		log.Tracef(format, args...)
	}
}

// Debugln - Logs the event at the Debug level
//
//goland:noinspection GoUnusedExportedFunction
func Debugln(args ...any) {
	if log.Level <= logrus.DebugLevel {
		log.Debugln(args...)
	}
}

// Debugf - Logs the event at the Debug level
//
//goland:noinspection GoUnusedExportedFunction
func Debugf(format string, args ...any) {
	if log.Level <= logrus.DebugLevel {
		log.Debugf(format, args...)
	}
}

// Infoln - Logs the event at the Infoln level
//
//goland:noinspection GoUnusedExportedFunction
func Infoln(args ...any) {
	if log.Level <= logrus.InfoLevel {
		log.Infoln(args...)
	}
}

// Warnln - Logs the event at the Warning level
//
//goland:noinspection GoUnusedExportedFunction
func Warnln(args ...any) {
	if log.Level <= logrus.WarnLevel {
		log.Warningln(args...)
	}
}

// Warnf - Logs the event at the Warning level with formatting
//
//goland:noinspection GoUnusedExportedFunction
func Warnf(format string, args ...any) {
	if log.Level == logrus.WarnLevel {
		log.Warnf(format, args...)
	}
}

// Errorln - Logs the event at the Error level
//
//goland:noinspection GoUnusedExportedFunction
func Errorln(args ...any) {
	if log.Level <= logrus.ErrorLevel {
		log.Errorln(args...)
	}
}

// Fatalln - Logs the event at the Fatal level
//
//goland:noinspection GoUnusedExportedFunction
func Fatalln(args ...any) {
	if log.Level <= logrus.FatalLevel {
		log.Fatalln(args...)
	}
}

//goland:noinspection GoUnusedExportedFunction
func FuncName() string {
	pc, _, line, _ := runtime.Caller(1)
	return fmt.Sprintf("(%s:L%d)", runtime.FuncForPC(pc).Name(), line)
}
