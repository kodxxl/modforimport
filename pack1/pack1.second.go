package pack1

import(
	"log/slog"
)

func nop(a, b int) int {
	return a + b
}

func init() {
	slog.Info("pack", 1, "second")
}