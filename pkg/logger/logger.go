package logger

import (
	"log/slog"
	"os"
	"time"
)

func init() {
	logOpts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Match the key that we want
			if a.Key == slog.TimeKey {
				a.Key = "date" //rename to date
				a.Value = slog.StringValue(time.Now().UTC().Format(time.RFC3339))
			}
			return a
		},
	}
	txtLog := slog.NewTextHandler(os.Stdout, logOpts)

	logger := slog.New(txtLog)
	slog.SetDefault(logger)
}
