package mob_smssdk

//go:generate go-easy generate error-code --type ErrorCode
type ErrorCode int

const (
	// AppKey为空
	ErrorCodeAppKeyEmpty ErrorCode = 405
	// AppKey无效
	ErrorCodeAppKeyInValid ErrorCode = 406
	// 国家代码或手机号码为空
	ErrorCodeZoneOrPhoneEmpty ErrorCode = 456
	// 手机号码格式错误
	ErrorCodePhoneFormat ErrorCode = 457
	// 请求校验的验证码为空
	ErrorCodeVerifyCodeEmpty ErrorCode = 466
	// 请求校验验证码频繁（5分钟内同一个appkey的同一个号码最多只能校验三次）
	ErrorCodeRequestFrequently ErrorCode = 467
	// 验证码错误
	ErrorCodeVerifyCode ErrorCode = 468
	// 没有打开服务端验证开关
	ErrorCodeServerVerifyNotOpen ErrorCode = 474
)