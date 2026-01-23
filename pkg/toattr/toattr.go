package toattr

import (
	"log/slog"
)

func Err(err error) slog.Attr {
	var str string
	if err != nil {
		str = err.Error()
	}
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(str),
	}
}
