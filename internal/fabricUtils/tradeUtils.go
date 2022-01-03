package fabricUtils

import (
	"github.com/gin-gonic/gin"
	"medicineApp/pkg/warpper"
)

// QueryMedicineByCode 通过溯源码查询药品，这个函数待替换
//func QueryMedicineByCode(userId, org, contractName, medicineCode string) (*entity.MedicineOnChain, error) {
//	// 获取指定合约链码
//	contract, err := GetContract(userId, org, contractName)
//	if err != nil {
//		return nil, err
//	}
//	// 查询
//	res, err := contract.EvaluateTransaction("QueryMedicineByCode", medicineCode)
//	if err != nil {
//		return nil, err
//	}
//	var medicine entity.MedicineOnChain
//	err = json.Unmarshal(res, &medicine)
//	if err != nil {
//		return nil, err
//	}
//	return &medicine, nil
//}

// TradePropose 发起交易
func TradePropose(c *gin.Context) {
	// 查询药品信息
	//medicine, err := QueryMedicineByCode(userId, org, contractName, medicineCode)
	//if err != nil {
	//	return err
	//}

	//if medicine.TractionType != "可交易" {
	//	return fmt.Errorf("指定的药品不可交易！")
	//} else if medicine.Owner != userName {
	//	return fmt.Errorf("当前用户不是药品的拥有者，无权发起交易！")
	//}

	userId := c.Query("userId")
	userName := c.Query("userName")
	org := c.Query("org")
	contractName := c.Query("contractName")
	medicineCode := c.Query("medicineCode")
	receiverName := c.Query("receiverName")
	// 连接网络获取合约
	contract, err := GetContract(userId, org, contractName)
	if err != nil {
		warpper.ResError(c, err)
	}
	res, err := contract.SubmitTransaction("TradePropose", userName, receiverName, medicineCode)
	if err != nil {
		warpper.ResError(c, err)
	}
	warpper.ResSuccess(c, res)
}

func TradeReceive(c *gin.Context) {
	//medicine, err := QueryMedicineByCode(userId, org, contractName, medicineCode)
	//if err != nil {
	//	return err
	//}
	//if medicine.TractionType != "待接收" {
	//	return fmt.Errorf("指定的药品不在待接收状态！")
	//} else if medicine.Receiver != userName {
	//	return fmt.Errorf("当前用户不是药品的接收方，无权接受交易！")
	//}

	userId := c.Query("userId")
	userName := c.Query("userName")
	org := c.Query("org")
	contractName := c.Query("contractName")
	medicineCode := c.Query("medicineCode")
	// 连接网络获取合约
	contract, err := GetContract(userId, org, contractName)
	if err != nil {
		warpper.ResError(c, err)
	}
	res, err := contract.SubmitTransaction("TradeReceive", userName, medicineCode)
	if err != nil {
		warpper.ResError(c, err)
	}
	warpper.ResSuccess(c, res)
}
