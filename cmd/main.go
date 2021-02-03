package main

import (
	"context"
	"flag"
	"github.com/buzyakabarbuzyaka/base/kit/config"
	"github.com/buzyakabarbuzyaka/base/kit/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	GlobalConfig := config.ServiceConfig{}.FromFile()
	log := logger.Init(GlobalConfig.LoggerConfig)

	gin.SetMode(GlobalConfig.ServerConfig.Mode)
	logInterface := gin.LoggerWithWriter(log.Writer(), "/status", "/metrics", "/healthz")

	router := gin.New()
	router.Use(logInterface, gin.Recovery())

	router.Any("/shiiish", handler)

	srv := &http.Server{
		Addr:    ":" + GlobalConfig.ServerConfig.Port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Panic(err)
	}

	log.Infoln("Shutting down")
	os.Exit(0)
}

func handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"bibo": "bobo",
	})
}
