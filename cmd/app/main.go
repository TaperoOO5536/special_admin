package main

import (
	"context"
	"log"
	"time"

	"github.com/TaperoOO5536/special_admin/internal/app"
	"github.com/TaperoOO5536/special_admin/internal/config"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	config.LoadEnv()

	cfg := &app.Config{
		GrpcPort:     "8090",
		HttpPort:     "8091",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Dsn:          config.GetDsn(),
	}

	app := app.New(cfg)

	if err := app.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	password, _ := bcrypt.GenerateFromPassword([]byte("qwer1234"), bcrypt.DefaultCost)
	log.Println(string(password))
}