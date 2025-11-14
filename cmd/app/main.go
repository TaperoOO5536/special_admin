package main

import (
	"context"
	"log"
	"time"

	"github.com/TaperoOO5536/special_admin/internal/app"
	"github.com/TaperoOO5536/special_admin/pkg/env"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	env.LoadEnv()

	cfg := &app.Config{
		GrpcPort:     env.GetGRPCPort(),
		HttpPort:     env.GetHTTPPort(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Dsn:          env.GetDsn(),
	}

	app := app.New(cfg)

	if err := app.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	password, _ := bcrypt.GenerateFromPassword([]byte("qwer1234"), bcrypt.DefaultCost)
	log.Println(string(password))
}