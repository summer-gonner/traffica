package response

import "fmt"

var (
	SUCCESS            = Status{Code: "000000", Message: "查询成功"}
	ES_CONNECT_SUCCESS = Status{Code: "000001", Message: "Es连接成功"}
)

type Status struct {
	Code    string
	Message string
}

// 定义一个方法来打印状态信息
func (s Status) String() string {
	return fmt.Sprintf("%s (Code: %d)", s.Message, s.Code)
}
