package bootstrap

import (
	"github.com/fadhlinw/clean-gin/api/controllers"
	"github.com/fadhlinw/clean-gin/api/middlewares"
	"github.com/fadhlinw/clean-gin/api/routes"
	"github.com/fadhlinw/clean-gin/lib"
	"github.com/fadhlinw/clean-gin/repository"
	"github.com/fadhlinw/clean-gin/services"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	middlewares.Module,
	repository.Module,
)
