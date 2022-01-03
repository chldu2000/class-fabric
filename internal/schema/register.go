package schema

// RegisterDistributorReqBodySchema 注册结构体
type RegisterDistributorReqBodySchema struct {
	CompanyName string `json:"company_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
	License     string `json:"license_id" binding:"required"`
	LegalPerson string `json:"legal_person" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	GMPID       string `json:"GMPID" binding:"required"`
}

// RegisterDistributorResBodySchema 注册结构体
type RegisterDistributorResBodySchema struct {
	CompanyName string `json:"company_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
	License     string `json:"license_id" binding:"required"`
	LegalPerson string `json:"legal_person" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	State       int    `json:"state" binding:"required"`
	GMPID       string `json:"GMPID" binding:"required"`
}

// RegisterManufacturerReqBodySchema 注册结构体
type RegisterManufacturerReqBodySchema struct {
	CompanyName string `json:"company_name" binding:"required"`
	RegisterId  string `json:"register_id" binding:"required"`
	Password    string `json:"password" binding:"required"`
	License     string `json:"license_id" binding:"required"`
	LegalPerson string `json:"legal_person" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	GMPID       string `json:"GMPID" binding:"required"`
}

// RegisterManufacturerResBodySchema 注册结构体
type RegisterManufacturerResBodySchema struct {
	CompanyName string `json:"company_name" binding:"required"`
	RegisterId  string `json:"register_id" binding:"required"`
	Password    string `json:"password" binding:"required"`
	License     string `json:"license_id" binding:"required"`
	LegalPerson string `json:"legal_person" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	State       int    `json:"state" binding:"required"`
	GMPID       string `json:"GMPID" binding:"required"`
}

// RegisterConsumerReqBodySchema 注册结构体
type RegisterConsumerReqBodySchema struct {
	Name     string `json:"name" binding:"required"`
	IdNumber string `json:"id_number" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterConsumerResBodySchema 注册结构体
type RegisterConsumerResBodySchema struct {
	Name     string `json:"name" binding:"required"`
	IdNumber string `json:"id_number" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	State    int    `json:"state" binding:"required"`
}
