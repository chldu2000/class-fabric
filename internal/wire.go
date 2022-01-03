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
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGorm,
		controller.ControllerSet,
		model.ModelSet,
		dao.DaoSet,
		router.RouterSet,
		InitGinEngine,
		InjectorSet,
	)

	return nil, nil, nil
}
