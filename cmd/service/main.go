package main

import (
	"fmt"

	"github.com/hse-revizor/projects-service/internal/di"
	"github.com/hse-revizor/projects-service/internal/utils/config"
	"github.com/hse-revizor/projects-service/internal/utils/flags"
	"github.com/slipneff/gogger"
	"github.com/slipneff/gogger/log"
)

// @title           Projects Service API
// @version         1.0.0
// @description     This is the Projects Service API documentation
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
// @host           localhost:8787
// @BasePath       /api
func main() {
	flags := flags.MustParseFlags()
	cfg := config.MustLoadConfig(flags.EnvMode)
	container := di.New(cfg)
	gogger.ConfigureZeroLogger()

	container.GetProjectService()
	log.Info(fmt.Sprintf("Server starting on %s:%d", cfg.Host, cfg.Port))

	err := container.GetHttpServer().ListenAndServe()
	if err != nil {
		log.Panic(err, "Fail serve HTTP")
	}
}
