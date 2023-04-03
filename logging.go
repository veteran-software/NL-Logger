package logging

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

//goland:noinspection GoUnusedConst
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
	DlYouTube   = "[DL-YT    ]"
	DlTwitch    = "[DL-TWITCH]"
	Stripe      = "[STRIPE   ]"
)

//goland:noinspection GoUnusedConst
const (
	PanicLevel = logrus.PanicLevel
	FatalLevel = logrus.FatalLevel
	ErrorLevel = logrus.ErrorLevel
	WarnLevel  = logrus.WarnLevel
	InfoLevel  = logrus.InfoLevel
	DebugLevel = logrus.DebugLevel
	TraceLevel = logrus.TraceLevel
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

// SetLevel - Changes the log level programmatically
//
//goland:noinspection GoUnusedExportedFunction
func SetLevel(level logrus.Level) {
	logrus.SetLevel(level)
}

// Traceln - Logs the event at the TraceLevel on a new line
//
//goland:noinspection GoUnusedExportedFunction
func Traceln(args ...any) {
	if log.IsLevelEnabled(TraceLevel) {
		log.Traceln(args...)
	}
}

// Tracef - Logs the event at the TraceLevel level with formatting
//
//goland:noinspection GoUnusedExportedFunction
func Tracef(format string, args ...any) {
	if log.IsLevelEnabled(TraceLevel) {
		log.Tracef(format, args...)
	}
}

// Trace - Logs the event at the TraceLevel
//
//goland:noinspection GoUnusedExportedFunction
func Trace(args ...any) {
	if log.IsLevelEnabled(TraceLevel) {
		log.Trace(args...)
	}
}

// Debugln - Logs the event at the DebugLevel
//
//goland:noinspection GoUnusedExportedFunction
func Debugln(args ...any) {
	if log.IsLevelEnabled(DebugLevel) {
		log.Debugln(args...)
	}
}

// Debugf - Logs the event at the DebugLevel
//
//goland:noinspection GoUnusedExportedFunction
func Debugf(format string, args ...any) {
	if log.IsLevelEnabled(DebugLevel) {
		log.Debugf(format, args...)
	}
}

// Infoln - Logs the event at the InfoLevel level
//
//goland:noinspection GoUnusedExportedFunction
func Infoln(args ...any) {
	if log.IsLevelEnabled(InfoLevel) {
		log.Infoln(args...)
	}
}

// Infof - Logs the event at the InfoLevel
//
//goland:noinspection GoUnusedExportedFunction
func Infof(format string, args ...any) {
	if log.IsLevelEnabled(InfoLevel) {
		log.Infof(format, args...)
	}
}

// Warnln - Logs the event at the WarnLevel
//
//goland:noinspection GoUnusedExportedFunction
func Warnln(args ...any) {
	if log.IsLevelEnabled(WarnLevel) {
		log.Warningln(args...)
	}
}

// Warnf - Logs the event at the WarnLevel with formatting
//
//goland:noinspection GoUnusedExportedFunction
func Warnf(format string, args ...any) {
	if log.IsLevelEnabled(WarnLevel) {
		log.Warnf(format, args...)
	}
}

// Errorln - Logs the event at the ErrorLevel
//
//goland:noinspection GoUnusedExportedFunction
func Errorln(args ...any) {
	if log.IsLevelEnabled(ErrorLevel) {
		log.Errorln(args...)
	}
}

// Errorf - Logs the event at the ErrorLevel with formatting
//
//goland:noinspection GoUnusedExportedFunction
func Errorf(format string, args ...any) {
	if log.IsLevelEnabled(ErrorLevel) {
		log.Errorf(format, args...)
	}
}

// Fatalln - Logs the event at the FatalLevel
//
//goland:noinspection GoUnusedExportedFunction
func Fatalln(args ...any) {
	if log.IsLevelEnabled(FatalLevel) {
		log.Fatalln(args...)
	}
}

// FuncName - Displays the file name and line number in logs
//
//goland:noinspection GoUnusedExportedFunction
func FuncName() string {
	pc, _, line, _ := runtime.Caller(1)
	return fmt.Sprintf("(%s:L%d)", runtime.FuncForPC(pc).Name(), line)
}
