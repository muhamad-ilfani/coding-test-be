package main

import (
	"coding-test-be/internal/app"
	"context"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app.Run(context.Background())
}
