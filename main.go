package main

import (
	"context"
	"flag"
	"github.com/gateway-server-go/config"
	"github.com/gateway-server-go/models"
	routes "github.com/gateway-server-go/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"time"
)



func StartServer() {
	var addr = flag.String("addr", "127.0.0.1:8000", "Http service address")
	flag.Parse()
	// Echo instance
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.INFO) // HTTP log

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("2M"))
	e.Use(middleware.CORS())

	// Request body validator
	// e.Validator = &structValidator{validator: validator.New()}

	// Routes
	routes.RegisterRestAPI(e)

	// Start server
	go func() {
		if err := e.Start(*addr); err != nil {
			e.Logger.Error(err)
			e.Logger.Info("An error occurred, shutting down the server...")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func main() {
	config.Init("./")
	models.InitDB()
	StartServer()
}