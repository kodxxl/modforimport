package pack1

import(
	"log/slog"
)

func Nop(){
	return
}

func init() {
	slog.Info("pack", 1, "second")
}