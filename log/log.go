package log

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
)

type log struct {
	*slog.Logger
}

var L log

func InitDefaultLog(app string) {
	L = log{
		Logger: slog.New(),
	}

	h1 := handler.MustFileHandler("/tmp/"+app+"_error.log", true)
	h1.Levels = slog.Levels{slog.PanicLevel, slog.ErrorLevel, slog.WarnLevel}

	h2 := handler.MustFileHandler("/tmp/"+app+"info.log", true)
	h2.Levels = slog.Levels{slog.InfoLevel, slog.NoticeLevel, slog.DebugLevel, slog.TraceLevel}

	h3 := handler.NewConsole(slog.AllLevels)

	L.AddHandlers(h1)
	L.AddHandler(h2)
	L.AddHandler(h3)
}
