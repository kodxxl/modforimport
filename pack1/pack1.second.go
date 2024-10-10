package pack1

import(
	"log/slog"
)

func nop() {
	return 42
}

func init() {
	slog.Info("pack", 1, "second")
}