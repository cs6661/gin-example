package dto

type MyCode int64

const (
	CodeInvalidPassword MyCode = 1004 // 密码错误
)

// 返回1xxxx -> 请求成功
const (
	Success          MyCode = 10000 // 请求成功的时候的返回值
	DataEmpty        MyCode = 10001 // 请求的数据为空时候的返回值
	ParameterMiss    MyCode = 10002 // 请求所需的参数部分缺失
	ParameterInvalid MyCode = 10003 // 请求的参数存在不合法的情况
	OpenIDEmpty      MyCode = 10004 // OpenID 为空
	CodeServerBusy   MyCode = 10005
)

// 返回2xxxx -> MySQL错误
const (
	MySQLQueryError     MyCode = 20000 // MySQL的查询错误
	MySQLScanError      MyCode = 20001 // MySQL在Scan时候的错误
	MySQLInsertError    MyCode = 20002 // MySQL的插入错误
	MySQLDeleteError    MyCode = 20003 // MySQL的删除错误
	MySQLUpdateError    MyCode = 20004 // MySQL的更新错误
	MySQLCountError     MyCode = 20004 // MySQL的计数错误
	MySQLAffectRowError MyCode = 20005 // MySQL获取影响行数错误
	MySQLPrepareError   MyCode = 20006
	MySQLExecError      MyCode = 20007
)

// 返回3xxxx -> 其他类型的错误
const (
	StrconvAtoiError MyCode = 30000 // 类型转换错误（主要是string转int）
	ParseTimeError   MyCode = 30001 // 日期类型的错误
	ReadFileError    MyCode = 30002 // 读取文件错误
	UnknowError      MyCode = 30003 // 未知错误
	UnknowOwner      MyCode = 30004 // 不是球馆管理员返回码
	InfoNotMatch     MyCode = 30005 // 信息不匹配
	ModelIsExist     MyCode = 30006 // 该类型已存在
	NotPermission    MyCode = 30007 // 用户无权限
)

// 返回4xxxx -> json错误
const (
	JsonUnmarshalError MyCode = 40000
	JsonMarshalError   MyCode = 40001
)

// 返回8xxxx -> HTTP错误
const (
	HTTPDoReqError  MyCode = 80000
	HTTPNewReqError MyCode = 80001
	HTTPGetError    MyCode = 80002
	HTTPPostError   MyCode = 80003
)

const (
	BizError MyCode = 90000
)

var msgFlags = map[MyCode]string{

	CodeServerBusy:      "服务繁忙",
	Success:             "请求成功",
	DataEmpty:           "请求的数据为空",
	ParameterMiss:       "请求所需的参数部分缺失",
	ParameterInvalid:    "请求的参数存在不合法的情况",
	OpenIDEmpty:         "OpenID 为空",
	CodeInvalidPassword: "账号或密码错误",

	MySQLQueryError:  "MySQL的查询错误",
	MySQLScanError:   "MySQL在Scan时候的错误",
	MySQLInsertError: "MySQL的插入错误",
	MySQLDeleteError: "MySQL的删除错误",
	MySQLUpdateError: "MySQL的更新错误",
	//MySQLCountError:     "MySQL的计数错误",
	MySQLAffectRowError: "MySQL获取影响行数错误",
	MySQLPrepareError:   "MySQL语句错误",
	MySQLExecError:      "MySQL执行错误",
	NotPermission:       "用户无权限",
}

func (c MyCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeServerBusy]
}
