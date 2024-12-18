package main

import "log/slog"

func main() {
	logger := slog.New(slog.Default().Handler())
	logger.Info("Hello, Arenko!")
}
