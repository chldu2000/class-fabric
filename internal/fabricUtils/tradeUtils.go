package fabricUtils

import (
	"github.com/gin-gonic/gin"
	"medicineApp/internal/schema"
	"medicineApp/pkg/warpper"
)

// TradePropose 发起交易
func TradePropose(c *gin.Context) {
	var data schema.TradeProposeRequestBody
	if err := warpper.ParseQuery(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	// 连接网络获取合约
	contract, err := GetContract(data.UserId, data.Org, data.ContractName)
	if err != nil {
		warpper.ResError(c, err)
	}
	res, err := contract.SubmitTransaction("TradePropose", data.UserName, data.ReceiverName, data.MedicineCode)
	if err != nil {
		warpper.ResError(c, err)
	}
	warpper.ResSuccess(c, res)
}

func TradeReceive(c *gin.Context) {
	var data schema.TradeReceiveRequestBody
	if err := warpper.ParseQuery(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}
	// 连接网络获取合约
	contract, err := GetContract(data.UserId, data.Org, data.ContractName)
	if err != nil {
		warpper.ResError(c, err)
	}
	res, err := contract.SubmitTransaction("TradeReceive", data.UserName, data.MedicineCode)
	if err != nil {
		warpper.ResError(c, err)
	}
	warpper.ResSuccess(c, res)
}
