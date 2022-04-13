package config

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/NinaLeven/TopSecretProject/internal/api"
	"github.com/NinaLeven/TopSecretProject/internal/integrations/employee"
	"github.com/NinaLeven/TopSecretProject/internal/integrations/storage"
	"github.com/NinaLeven/TopSecretProject/internal/projectmanager"
)

type ServiceConfig struct {
	API              api.Config   `yaml:"api"`
	PG               string       `yaml:"pg"`
	LogLevel         logrus.Level `yaml:"log_level"`
	EmployeesBaseURL string       `yaml:"employees_base_url"`
}

func Configure(ctx context.Context, configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("unable to open config: %w", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	config := ServiceConfig{}

	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return fmt.Errorf("unable to unmarshal config: %w", err)
	}

	log := &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.JSONFormatter{},
		Level:     config.LogLevel,
	}

	db, err := sqlx.ConnectContext(ctx, "postgres", config.PG)
	if err != nil {
		return fmt.Errorf("unable to connect to db: %w", err)
	}

	storageFactory := storage.NewFactory(db, log)

	employeeSevice, err := employee.NewEmployeeServiceClient(&http.Client{}, config.EmployeesBaseURL)
	if err != nil {
		return fmt.Errorf("unable to create employee clisnt: %w", err)
	}

	projectManagementService := projectmanager.NewProjectManagementService(employeeSevice)

	projectManagerServer := api.NewProjectManagerServer(&config.API, log, projectManagementService, storageFactory)

	router := chi.NewRouter()
	projectManagerServer.Route(router)

	l, err := net.Listen("tcp", config.API.Addr)
	if err != nil {
		return fmt.Errorf("unable to listen to addr: %w", err)
	}

	httpServer := &http.Server{
		Addr:    config.API.Addr,
		Handler: router,
	}

	go func() {
		<-ctx.Done()
		cctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		err := httpServer.Shutdown(cctx)
		if err != nil {
			log.WithError(err).Error("unable to shutdown server")
		}
	}()

	log.Info("starting server")

	err = httpServer.Serve(l)
	if err != nil {
		return fmt.Errorf("server error: %w", err)
	}

	log.Info("server exited")

	return nil
}
