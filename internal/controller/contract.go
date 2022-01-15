package controller

import (
	"medicineApp/internal/schema"
	"medicineApp/internal/service"
	"medicineApp/pkg/warpper"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ContractControllerSet = wire.NewSet(wire.Struct(new(ContractController), "*"))

type ContractController struct {
	ContractService *service.MedicineContract
}

func (cc *ContractController) Add(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.Medicine
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}
	err := cc.ContractService.Add(ctx, data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}
	warpper.ResSuccess(c, nil)
}
func (cc *ContractController) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.Medicine
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}
	err := cc.ContractService.Delete(ctx, data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}
	warpper.ResSuccess(c, nil)
}
func (cc *ContractController) Update(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.Medicine
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}
	err := cc.ContractService.Update(ctx, data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}
	warpper.ResSuccess(c, nil)
}

func (cc *ContractController) Query(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.QueryMedicine
	if err := warpper.ParseQuery(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}
	res, err := cc.ContractService.Query(ctx, data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}
	warpper.ResSuccess(c, res)
}

//func (cc *ContractController) QueryHistoryByCode(c *gin.Context) {
//	ctx := c.Request.Context()
//
//	var data schema.QueryMedicineHistory
//	if err := warpper.ParseQuery(c, &data); err != nil {
//		warpper.ResError(c, err)
//		return
//	}
//	res, err := cc.ContractService.
//
//}

