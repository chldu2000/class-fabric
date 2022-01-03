package schema

// LoginReqBodySchema 登录结构体
type LoginReqBodySchema struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginAdminReqBodySchema 登录结构体
type LoginAdminReqBodySchema struct {
	SystemAccount string `json:"system_account" binding:"required"`
	Password      string `json:"password" binding:"required"`
}

// LoginConsumerReqBodySchema 登录结构体
type LoginConsumerReqBodySchema struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginCompanyReqBodySchema 登录结构体
type LoginCompanyReqBodySchema struct {
	CompanyName string `json:"company_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

// LoginResBodySchema 需要返回给前端的数据
type LoginResBodySchema struct {
	UID       string `json:"uid"`
	Authority int    `json:"authority"`
}

// LoginAdminResBodySchema 需要返回给前端的数据
type LoginAdminResBodySchema struct {
	SystemAccount string `json:"system_account"`
	Token         string `json:"token"`
}

// LoginConsumerResBodySchema 需要返回给前端的数据
type LoginConsumerResBodySchema struct {
	Name     string `json:"name"`
	IdNumber string `json:"id_number"`
	Token    string `json:"token"`
}

// LoginCompanyResBodySchema 需要返回给前端的数据
type LoginCompanyResBodySchema struct {
	CompanyName string `json:"company_name"`
	Token       string `json:"token"`
}
