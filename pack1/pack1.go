package pack1

import( "log/slog" )

func Add(x, y int) int {

  return x + y
}

func init() {
	slog.Info("pack", 1, "first")
}