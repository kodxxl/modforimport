package pack2

import(
	"log/slog"
)

func Nop(a, b int) int {
	return a + b
}

func nop2() int {
	return 42
}

func init(){
	slog.Info("pack", 2, "first")
}