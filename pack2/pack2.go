package pack2

import(
	"log/slog"
)

func nop() int {
	return 42
}

func nop2() int {
	return 42
}

func init(){
	slog.Info("pack", 2, "first")
}