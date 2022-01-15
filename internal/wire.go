//go:generate wire
//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
	"medicineApp/internal/controller"
	"medicineApp/internal/dao"
	"medicineApp/internal/model"
	"medicineApp/internal/router"
	"medicineApp/internal/service"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGorm,
		controller.ControllerSet,
		model.ModelSet,
		dao.DaoSet,
		service.ServiceSet,
		router.RouterSet,
		InitGinEngine,
		InjectorSet,
	)

	return nil, nil, nil
}
