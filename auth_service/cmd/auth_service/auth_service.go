package main

import (
	"github.com/lallison21/call_center_request/auth_service/internal/application"
	"github.com/lallison21/call_center_request/auth_service/internal/config"
)

func main() {
	cfg := config.New()

	app := application.New(cfg)

	app.Run()
}
