package common

// 自定义业务错误
type BizError struct {
	Code int
	Msg  string
}

func (err *BizError) Error() string {
	return err.Msg
}

func New(code int, msg string) *BizError {
	return &BizError{
		Code: code,
		Msg:  msg,
	}
}

var (
	ParamError    = New(400, "参数错误")
	StockError    = New(5001, "库存不足")
	InternalError = New(500, "服务端错误")
)
