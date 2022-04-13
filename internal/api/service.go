package api

import (
	"encoding/json"
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"github.com/NinaLeven/TopSecretProject/internal/integrations/storage"
	"github.com/NinaLeven/TopSecretProject/internal/projectmanager"
)

type ProjectManagerServer interface {
	Route(router chi.Router)
}

type Config struct {
	BasePath string `yaml:"base_path"`
	Addr     string `yaml:"addr"`
}

type projectManagerServer struct {
	config *Config

	projectManagementService projectmanager.ProjectManagementService
	storageFactory           storage.Factory

	log *logrus.Logger
}

func NewProjectManagerServer(
	config *Config,
	log *logrus.Logger,
	projectManagementService projectmanager.ProjectManagementService,
	storageFactory storage.Factory,
) ProjectManagerServer {
	return &projectManagerServer{
		config:                   config,
		log:                      log,
		projectManagementService: projectManagementService,
		storageFactory:           storageFactory,
	}
}

func (p *projectManagerServer) Route(router chi.Router) {
	router.Route(path.Join(p.config.BasePath, "api", "projects"), func(router chi.Router) {
		router.Get("/", p.listProjects)
		router.Post("/", p.createProject)
		router.Get("/{uid}", p.getProject)
		router.Patch("/{uid}", p.updateProject)
	})
}

type requestIDContextKey struct{}

func respondWithJSON(w http.ResponseWriter, code int, err interface{}) {
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(err)
}
