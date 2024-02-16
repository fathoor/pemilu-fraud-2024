package main

import (
	"github.com/fathoor/pemilu-fraud-2024/internal/config"
	"github.com/fathoor/pemilu-fraud-2024/internal/exception"
	"github.com/fathoor/pemilu-fraud-2024/internal/provider"
)

func main() {
	var app = config.ProvideApp()

	provider.ProvideModule(app)

	err := app.Listen(":2024")
	exception.PanicIfNeeded(err)
}
