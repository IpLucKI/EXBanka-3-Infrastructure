package main

import (
	"log/slog"
	"os"

	"github.com/RAF-SI-2025/EXBanka-3-Infrastructure/internal/config"
	"github.com/RAF-SI-2025/EXBanka-3-Infrastructure/internal/database"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		slog.Error("Failed to connect", "error", err)
		os.Exit(1)
	}

	if err := database.Migrate(db); err != nil {
		slog.Error("Migration failed", "error", err)
		os.Exit(1)
	}

	if err := database.SeedPermissions(db); err != nil {
		slog.Error("Failed to seed permissions", "error", err)
		os.Exit(1)
	}

	if err := database.SeedDefaultAdmin(db); err != nil {
		slog.Error("Failed to seed default admin", "error", err)
		os.Exit(1)
	}

	slog.Info("Admin user created successfully", "email", "admin@bank.com", "password", "Admin123!")
}
