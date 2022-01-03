package schema

// TestHandlerResBodySchema 需要返回给前端的数据
type TestHandlerResBodySchema struct {
	Name     string `json:"name"`
	IdNumber string `json:"id_number"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
