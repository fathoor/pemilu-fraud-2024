package main

import (
	"fmt"
	"github.com/fathoor/fraud24/internal/config"
	"github.com/fathoor/fraud24/internal/exception"
	"github.com/fathoor/fraud24/internal/provider"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var (
		cfg = config.ProvideConfig()
		app = config.ProvideApp()
	)

	provider.ProvideModule(app)

	err := app.Listen(fmt.Sprintf(":%s", cfg.Get("APP_PORT")))
	exception.PanicIfNeeded(err)
}
