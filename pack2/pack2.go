package pack2

import(
	"log/slog"
)

func nop() {

}

func nop2() {

}

func init(){
	slog.Info("pack", 2, "first")
}