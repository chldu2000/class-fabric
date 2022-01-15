package fabricUtils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"medicineApp/internal/schema"
	"medicineApp/pkg/warpper"
)

func QueryMedicineHistoryByCode(c *gin.Context) {
	var data schema.QueryMedicineHistory
	if err := warpper.ParseQuery(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	// 连接网络获取合约
	contract, err := GetContract(data.UserId, data.Org, data.ContractName)
	if err != nil {
		warpper.ResError(c, err)
	}
	res, err := contract.EvaluateTransaction("QueryMedicineHistoryByCode", data.MedicineCode)
	if err != nil {
		warpper.ResError(c, err)
		log.Println("获取不到药品交易历史")
	}
	medicineHistory := json.RawMessage(res)
	warpper.ResSuccess(c, medicineHistory)
}
