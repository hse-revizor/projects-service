package di

import (
	"fmt"
	"net/http"
	"time"

	trmgorm "github.com/avito-tech/go-transaction-manager/gorm"
	"github.com/avito-tech/go-transaction-manager/trm"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"github.com/hse-revizor/projects-service/internal/pkg/router"
	"github.com/hse-revizor/projects-service/internal/pkg/service/project"
	"github.com/hse-revizor/projects-service/internal/pkg/storage/sql"
	"github.com/hse-revizor/projects-service/internal/utils/config"
	"gorm.io/gorm"
)

type Container struct {
	cfg                *config.Config
	handler            *router.Handler
	storage            *sql.Storage
	httpServer         *http.Server
	db                 *gorm.DB
	transactionManager trm.Manager
	projectService     *project.Service
}

func New(cfg *config.Config) *Container {
	return &Container{cfg: cfg}
}

func (c *Container) GetDB() *gorm.DB {
	return get(&c.db, func() *gorm.DB {
		return sql.MustNewPostgresDB(c.cfg)
	})
}

func (c *Container) GetSQLStorage() *sql.Storage {
	return get(&c.storage, func() *sql.Storage {
		return sql.New(c.GetDB(), trmgorm.DefaultCtxGetter)
	})
}
func (c *Container) GetHttpHandler() *router.Handler {
	return get(&c.handler, func() *router.Handler {
		return router.NewRouter(c.cfg, c.GetProjectService())
	})
}
func (c *Container) GetHttpServer() *http.Server {
	return get(&c.httpServer, func() *http.Server {
		return &http.Server{
			Addr:           fmt.Sprintf("%s:%d", c.cfg.Host, c.cfg.Port),
			Handler:        c.GetHttpHandler().InitRoutes(),
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		}
	})
}

func (c *Container) GetTransactionManager() trm.Manager {
	return get(&c.transactionManager, func() trm.Manager {
		return manager.Must(trmgorm.NewDefaultFactory(c.GetDB()))
	})
}

func (c *Container) GetProjectService() *project.Service {
	return get(&c.projectService, func() *project.Service {
		return project.New(c.GetSQLStorage())
	})
}

func get[T comparable](obj *T, builder func() T) T {
	if *obj != *new(T) {
		return *obj
	}

	*obj = builder()
	return *obj
}
