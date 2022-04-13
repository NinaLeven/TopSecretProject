package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"

	"github.com/NinaLeven/TopSecretProject/internal/config"
)

func main() {
	ctx := context.Background()

	ctx, cancel := signal.NotifyContext(ctx, os.Kill, os.Interrupt)
	defer cancel()

	configPath := flag.String("c", "../..config/config.yaml", "path to config file")
	flag.Parse()

	if configPath == nil {
		fmt.Println("config path is not set")
		os.Exit(1)
	}

	log := &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.JSONFormatter{},
		Level:     logrus.ErrorLevel,
	}

	err := config.Configure(ctx, *configPath)
	if err != nil {
		cancel()
		log.WithError(err).Error("service exited")
		return
	}
}
