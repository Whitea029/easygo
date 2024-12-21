package api

import (
	"{{ .ProjectName }}/api/middlewares"
	"{{ .ProjectName }}/api/routers"
	"{{ .ProjectName }}/conf"
	"{{ .ProjectName }}/pkg/logging"
	"fmt"

	"github.com/gin-gonic/gin"
)

var logger = logging.GetLogger()
var ServerPort = conf.GetConfig().Server.Port

func InitServer() {
	r := gin.New()
	r.Use(gin.Logger())

	RegisterValidator()
	RegisterMiddleware(r)
	RegisterRouter(r)

	logger.Info(logging.General, logging.Startup, "Started", nil)
	err := r.Run(fmt.Sprintf(":%s", ServerPort))
	if err != nil {
		logger.Fatal(logging.General, logging.Startup, err.Error(), nil)
	}
}

func RegisterRouter(r *gin.Engine) {

	api := r.Group("/api")
	v1 := api.Group("/v1")

	{
		health := v1.Group("/health")
		routers.Health(health)
	}
}

func RegisterMiddleware(r *gin.Engine) {
	r.Use(gin.CustomRecovery(middlewares.ErrorHandler))
}

func RegisterValidator() {}
