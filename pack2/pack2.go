package pack2

import(
	"log/slog"
)

func nop() {
	return 42
}

func nop2() {
	return 42
}

func init(){
	slog.Info("pack", 2, "first")
}