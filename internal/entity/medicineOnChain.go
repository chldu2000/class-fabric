package entity

type MedicineOnChain struct {
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
