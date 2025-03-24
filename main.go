package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lenna-ai/azureOneSmile.git/config"
	appconfig "github.com/lenna-ai/azureOneSmile.git/config/appConfig"
	"github.com/lenna-ai/azureOneSmile.git/routes"
)

func main() {
	app := fiber.New()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	serverShutdown := make(chan struct{})

	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
		serverShutdown <- struct{}{}
	}()
	appconfig.InitApplication()
	// app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			// Izinkan semua subdomain dari bni.co.id
			return strings.HasSuffix(origin, ".bni.co.id") || origin == "http://bni.co.id"
		},
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	config.Logger(app)
	routes.Router(app)
	if err := app.Listen(":3000"); err != nil {
		panic(err.Error())
	}

	<-serverShutdown

	config.GeneralLogger.Println("Running cleanup tasks...")
	fmt.Println("Running cleanup tasks...")
}
