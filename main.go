package main

import "log/slog"

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		return
	}
	slog.Info("All systems offline")
}

func run() error {

	return nil
}
