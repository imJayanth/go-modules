package app

import (
	"log"
	"strconv"

	"github.com/imJayanth/go-modules/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func StartApplication(appConfig *config.AppConfig, mapUrls func(fApp *fiber.App, appConfig *config.AppConfig)) {
	var app = fiber.New()
	addLogger(appConfig, app)
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(recover.New())
	mapUrls(app, appConfig)
	log.Fatal(app.Listen(":" + strconv.Itoa(appConfig.ServerConfig.APIPort)))
}

func addLogger(appConfig *config.AppConfig, app *fiber.App) {
	if appConfig.LoggerConfig.File != nil {
		app.Use(logger.New(logger.Config{
			Output: appConfig.LoggerConfig.File,
		}))
	} else {
		app.Use(logger.New())
	}
}
