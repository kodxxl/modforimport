package pack1

import(
	"log/slog"
)

func nop() {

}

func init() {
	slog.Info("pack", 1, "second")
}