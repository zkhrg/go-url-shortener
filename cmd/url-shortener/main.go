package main

import (
	"fmt"
	"go-url-shortener/internal/config"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init storage: sqlite3

	// TODO: init router: chi, "chi render"

	// TODO: run server

}
