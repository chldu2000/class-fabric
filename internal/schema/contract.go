package schema

type Medicine struct {
	MedicineName  string `json:"medicineName"`
	ApprovalNo    string `json:"approval_no"`
	Spacification string `json:"spacification"`
	ProduceDate   string `json:"produceDate"`
	Producer      int    `json:"producer"`
	Status        string `json:"status"`
	BatchNo       string `json:"batchNo"`
	Expiration    string `json:"expiration"`
	Num           int    `json:"num"`
	MedicineCode  string `json:"medicineCode"`
	MedicineId    int    `json:"medicineId"`
	Unit          string `json:"unit"`
	CreateDate    string `json:"creatDate "`
	Owner         string `json:"owner"`
	Receiver      string `json:"receiver"`
	TractionType  string `json:"tractionType"`
}

type QueryMedicineRes struct {
	MedicineName  string `json:"medicineName"`
	ApprovalNo    string `json:"approval_no"`
	Spacification string `json:"spacification"`
	ProduceDate   string `json:"produceDate"`
	Producer      string `json:"producer"`
	Status        string `json:"status"`
	BatchNo       string `json:"batchNo"`
	Expiration    string `json:"expiration"`
	Num           int    `json:"num"`
	MedicineId    int    `json:"medicineId"`
}

type QueryInterface struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

type CommonReq struct {
	PageSize int `form:"pageSize"`
	Offset   int `form:"offset"`
}

type QueryMedicine struct {
	CommonReq
	Name     string `form:"name"`
	Producer string `form:"producer"`
	Approval string `form:"approval"`
}

type QueryMedicineHistory struct {
	CommonReq
	UserId string `form:"user_id"`
	Org string `form:"org"`
	ContractName string `form:"contract_name"`
	MedicineCode string `form:"medicine_code"`
}

type TradeProposeRequestBody struct {
	CommonReq
	UserId string `form:"user_id"`
	UserName string `form:"user_name"`
	Org string `form:"org"`
	ContractName string `form:"contract_name"`
	MedicineCode string `form:"medicine_code"`
	ReceiverName string `form:"receiver_name"`
}

type TradeReceiveRequestBody struct {
	CommonReq
	UserId string `form:"user_id"`
	UserName string `form:"user_name"`
	Org string `form:"org"`
	ContractName string `form:"contract_name"`
	MedicineCode string `form:"medicine_code"`
}
