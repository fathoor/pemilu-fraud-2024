package main

import (
	"github.com/fathoor/fraud24/internal/config"
	"github.com/fathoor/fraud24/internal/exception"
	"github.com/fathoor/fraud24/internal/provider"
)

func main() {
	var app = config.ProvideApp()

	provider.ProvideModule(app)

	err := app.Listen(":2024")
	exception.PanicIfNeeded(err)
}
