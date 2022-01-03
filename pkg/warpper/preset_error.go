package warpper

// 定义错误
var (
	ErrBadRequest              = New400Response("请求发生错误")
	ErrInvalidParent           = New400Response("无效的父级节点")
	ErrNotAllowDeleteWithChild = New400Response("含有子级，不能删除")
	ErrNotAllowDelete          = New400Response("资源不允许删除")
	ErrInvalidUserName         = New400Response("无效的用户名")
	ErrInvalidPassword         = New400Response("无效的密码")
	ErrInvalidUser             = New400Response("无效的用户")
	ErrCanNotGenerateToken     = New400Response("生成token失败")
	ErrUserDisable             = New400Response("用户被禁用，请联系管理员")
	ErrWrongCaptcha            = New400Response("验证码错误")
	ErrExpiredCaptcha          = New400Response("验证码过期")

	ErrNoAuth          = NewResponse(402, 402, "权限不够修改成员")
	ErrNoPerm          = NewResponse(401, 401, "无访问权限")
	ErrInvalidToken    = NewResponse(9999, 401, "令牌失效")
	ErrInvalidState    = NewResponse(9998, 200, "未通过注册")
	ErrNotFound        = NewResponse(404, 404, "资源不存在")
	ErrMethodNotAllow  = NewResponse(405, 405, "方法不被允许")
	ErrTooManyRequests = NewResponse(429, 429, "请求过于频繁")
	ErrInternalServer  = NewResponse(500, 500, "服务器发生错误")
)
