package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/wire"
	"io/ioutil"
	"medicineApp/internal/dao"
	"medicineApp/internal/schema"
	"os/exec"
	"strconv"
	"time"
)

type MedicineChain struct {
	MedicineCode string `json:"medicineCode"` // 药品防伪码
	MedicineName string `json:"medicineName"` // 药品名称
	MedicineId   string `json:"medicineId"`   // 药品Id
	ApprovalNo   string `json:"approvalNo"`   // 药品批准文号
	Unit         string `json:"unit"`         // 包装单位
	CreateDate   string `json:"createDate"`   // 上链时间
	Owner        string `json:"owner"`        // 所有者
	Receiver     string `json:"receiver"`     // 待接收交易者
	TractionType string `json:"tractionType"` // 最新交易类型
}

var MedicineServiceSet = wire.NewSet(wire.Struct(new(MedicineContract), "*"))

type MedicineContract struct {
	MedicineDao *dao.MedicineDao
}

func ExecCommand(strCommand string) ([]byte, error) {
	cmd := exec.Command(strCommand)

	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	outBytes, _ := ioutil.ReadAll(stdout)
	err := stdout.Close()
	if err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	return outBytes, nil
}

func GenerateCommands(funcName string, args ...string) string {
	const baseCommand = "peer chaincode invoke -o localhost:9050 --ordererTLSHostnameOverride orderer.medicine.com" +
		" --tls --cafile ${HOME}/go/src/github.com/hyperledger/fabric-samples/test-network-new/organizations/ordererOrganizations/medicine.com/orderers/orderer.medicine.com/msp/tlscacerts/tlsca.medicine.com-cert.pem" +
		" -C mychannel -n trade --peerAddresses localhost:5051 --tlsRootCertFiles ${HOME}/go/src/github.com/hyperledger/fabric-samples/test-network-new/organizations/peerOrganizations/orgAgency.medicine.com/peers/peer0.orgAgency.medicine.com/tls/ca.crt" +
		" --peerAddresses localhost:6051 --tlsRootCertFiles ${HOME}/go/src/github.com/hyperledger/fabric-samples/test-network-new/organizations/peerOrganizations/orgConsumer.medicine.com/peers/peer0.orgConsumer.medicine.com/tls/ca.crt --peerAddresses localhost:7051" +
		" --tlsRootCertFiles ${HOME}/go/src/github.com/hyperledger/fabric-samples/test-network-new/organizations/peerOrganizations/orgManufacturer.medicine.com/peers/peer0.orgManufacturer.medicine.com/tls/ca.crt --peerAddresses localhost:8051" +
		" --tlsRootCertFiles ${HOME}/go/src/github.com/hyperledger/fabric-samples/test-network-new/organizations/peerOrganizations/orgSupplier.medicine.com/peers/peer0.orgSupplier.medicine.com/tls/ca.crt"

	count := len(args)
	if count != 2 {
		return (baseCommand + " -c '{\"Args\":[\"Delete\"]}'")
	}
	return baseCommand + fmt.Sprintf(" -c '{\"Args\":[\"%s\", \"%s\", \"%s\"]}'", funcName, args[0], args[1])
}

func (m *MedicineContract) Add(ctx context.Context, params schema.Medicine) error {
	res, err := m.MedicineDao.Add(ctx, params)
	if err != nil {
		return err
	}
	medicine := MedicineChain{
		MedicineCode: strconv.Itoa(int(res.ID)),
		MedicineName: res.MedicineName,
		MedicineId:   strconv.Itoa(int(res.ID)),
		ApprovalNo:   res.ApprovalNo,
		Unit:         params.Unit,
		CreateDate:   time.Now().Format("2006-01-02"),
		Owner:        params.Owner,
		Receiver:     params.Receiver,
		TractionType: params.TractionType,
	}
	medicineStr, err := json.Marshal(medicine)
	if err != nil {
		return err
	}

	command := GenerateCommands("Create", medicine.MedicineId, string(medicineStr))
	//command := fmt.Sprintf("peer chaincode invoke -n mycc -c '{\"Args\":[\"Create\", \"%s\", \"%s\"]}' -C myc",
	//	medicine.MedicineId,
	//	string(medicineStr),
	//)
	_, err = ExecCommand(command)
	if err != nil {
		return err
	}
	return nil
}

func (m *MedicineContract) Update(ctx context.Context, params schema.Medicine) error {
	medicine := MedicineChain{
		MedicineCode: strconv.Itoa(params.MedicineId),
		MedicineName: params.MedicineName,
		MedicineId:   strconv.Itoa(params.MedicineId),
		ApprovalNo:   params.ApprovalNo,
		Unit:         params.Unit,
		CreateDate:   params.CreateDate,
		Owner:        params.Owner,
		Receiver:     params.Receiver,
		TractionType: params.TractionType,
	}
	medicineStr, err := json.Marshal(medicine)
	if err != nil {
		return err
	}
	command := GenerateCommands("Update", medicine.MedicineId, string(medicineStr))
	_, err = ExecCommand(command)
	if err != nil {
		return err
	}
	err = m.MedicineDao.Update(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (m *MedicineContract) Delete(ctx context.Context, params schema.Medicine) error {
	command := GenerateCommands("Delete", string(rune(params.MedicineId)))
	_, err := ExecCommand(command)
	if err != nil {
		return err
	}
	err = m.MedicineDao.Delete(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (m *MedicineContract) Query(ctx context.Context, params schema.QueryMedicine) (*schema.QueryInterface, error) {
	res, count, err := m.MedicineDao.Query(ctx, params)
	if err != nil {
		return nil, err
	}
	var result []schema.QueryMedicineRes
	for _, item := range *res {
		tmp := schema.QueryMedicineRes{
			MedicineName:  item.MedicineName,
			ApprovalNo:    item.ApprovalNo,
			Spacification: item.Spacification,
			ProduceDate:   item.ProduceDate.Format("2006-01-02"),
			Producer:      item.Manufacturers.CompanyName,
			Status:        item.Status,
			BatchNo:       item.BatchNo,
			Expiration:    item.Expiration,
			Num:           item.Num,
			MedicineId:    int(item.ID),
		}
		result = append(result, tmp)
	}
	return &schema.QueryInterface{
		Total: count,
		Data:  result,
	}, err
}
//
//func QueryMedicineHistoryByCode(c *gin.Context) {
//	userId := c.Query("userId")
//	org := c.Query("org")
//	contractName := c.Query("contractName")
//	medicineCode := c.Query("medicineCode")
//	// 连接网络获取合约
//	contract, err := fabricUtils.GetContract(userId, org, contractName)
//	if err != nil {
//		warpper.ResError(c, err)
//	}
//	res, err := contract.EvaluateTransaction("QueryMedicineHistoryByCode", medicineCode)
//	if err != nil {
//		warpper.ResError(c, err)
//		log.Println("获取不到药品交易历史")
//	}
//	medicineHistory := json.RawMessage(res)
//	warpper.ResSuccess(c, medicineHistory)
//}
//
//// TradePropose 发起交易
//func TradePropose(c *gin.Context) {
//	userId := c.Query("userId")
//	userName := c.Query("userName")
//	org := c.Query("org")
//	contractName := c.Query("contractName")
//	medicineCode := c.Query("medicineCode")
//	receiverName := c.Query("receiverName")
//	// 连接网络获取合约
//	contract, err := fabricUtils.GetContract(userId, org, contractName)
//	if err != nil {
//		warpper.ResError(c, err)
//	}
//	res, err := contract.SubmitTransaction("TradePropose", userName, receiverName, medicineCode)
//	if err != nil {
//		warpper.ResError(c, err)
//	}
//	warpper.ResSuccess(c, res)
//}
//
//func TradeReceive(c *gin.Context) {
//	userId := c.Query("userId")
//	userName := c.Query("userName")
//	org := c.Query("org")
//	contractName := c.Query("contractName")
//	medicineCode := c.Query("medicineCode")
//	// 连接网络获取合约
//	contract, err := fabricUtils.GetContract(userId, org, contractName)
//	if err != nil {
//		warpper.ResError(c, err)
//	}
//	res, err := contract.SubmitTransaction("TradeReceive", userName, medicineCode)
//	if err != nil {
//		warpper.ResError(c, err)
//	}
//	warpper.ResSuccess(c, res)
//}
