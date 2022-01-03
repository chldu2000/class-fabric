package schema

// PassReqBodySchema 审核结构体
type PassReqBodySchema struct {
	Id    uint
	State int `json:"state" binding:"required"`
}

// PassConsumerResBodySchema 审核结构体
type PassConsumerResBodySchema struct {
	Name     string `json:"name" binding:"required"`
	IdNumber string `json:"id_number" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	State    int    `json:"state" binding:"required"`
	Wallet   string `json:"wallet" binding:"required"`
}

// PassManufacturerResBodySchema 审核结构体
type PassManufacturerResBodySchema struct {
	CompanyName string `json:"company_name" binding:"required"`
	RegisterId  string `json:"register_id" binding:"required"`
	License     string `json:"license_id" binding:"required"`
	LegalPerson string `json:"legal_person" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	State       int    `json:"state" binding:"required"`
	Ip          string `json:"ip" binding:"required"`
	GMPID       string `json:"GMPID" binding:"required"`
	Wallet      string `json:"wallet" binding:"required"`
}

// PassDistributorResBodySchema 审核结构体
type PassDistributorResBodySchema struct {
	CompanyName string `json:"company_name" binding:"required"`
	License     string `json:"license_id" binding:"required"`
	LegalPerson string `json:"legal_person" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	State       int    `json:"state" binding:"required"`
	Ip          string `json:"ip" binding:"required"`
	GMPID       string `json:"GMPID" binding:"required"`
	Wallet      string `json:"wallet" binding:"required"`
}
