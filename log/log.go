package log

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
)

type log struct {
	*slog.Logger
}

type LogConf struct {
	Path  string `yaml:"path"`
	Every string `yaml:"every"`
	Env   string `yaml:"env"`
}

func RotateMap() map[string]uint8 {
	return map[string]uint8{
		"day":  0,
		"hour": 1,
		"60m":  1,
		"30m":  2,
		"15m":  3,
		"m":    4,
	}
}

var L log

func InitDefaultLog(c *LogConf) {
	L = log{
		Logger: slog.New(),
	}
	h, err := handler.NewRotateFileHandler(c.Path, handler.EveryDay)
	if err != nil {
		panic(err)
	}
	if c.Env == "prod" {
		h.Levels = slog.DangerLevels
		L.AddHandlers(h)
	} else {
		// h.Levels = slog.AllLevels

		console := handler.NewConsole(slog.AllLevels)
		L.AddHandler(console)
	}

}
