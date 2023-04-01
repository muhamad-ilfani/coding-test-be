package main

import (
	"coding-test-be/app"
	"context"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app.Run(context.Background())
}
