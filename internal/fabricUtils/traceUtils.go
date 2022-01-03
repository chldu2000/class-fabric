package fabricUtils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"medicineApp/pkg/warpper"
)

func QueryMedicineHistoryByCode(c *gin.Context) {
	userId := c.Query("userId")
	org := c.Query("org")
	contractName := c.Query("contractName")
	medicineCode := c.Query("medicineCode")
	// 连接网络获取合约
	contract, err := GetContract(userId, org, contractName)
	if err != nil {
		warpper.ResError(c, err)
	}
	res, err := contract.EvaluateTransaction("QueryMedicineHistoryByCode", medicineCode)
	if err != nil {
		warpper.ResError(c, err)
		log.Println("获取不到药品交易历史")
	}
	medicineHistory := json.RawMessage(res)
	warpper.ResSuccess(c, medicineHistory)
}
