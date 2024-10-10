package pack2

import(
	"log/slog"
)

func Nop(){
	return
}

func nop2() int {
	return 42
}

func init(){
	slog.Info("pack", 2, "first")
}