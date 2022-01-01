package errorcodes

import (
	. "github.com/bondhan/godddnews/infrastructure/client"
	"net/http"
)

const (
	ERR40001 ErrorCode = 40001
	ERR40002 ErrorCode = 40002
	ERR40003 ErrorCode = 40003
	ERR40004 ErrorCode = 40004
	ERR40005 ErrorCode = 40005
)

var ErrorCodesPayment = map[ErrorCode]ErrorMessageDict{
	ERR40001: {Status: http.StatusBadRequest, Type: "DATA_INVALID", Code: ERR40001, Message: "Invalid Data Request"},
	ERR40002: {Status: http.StatusInternalServerError, Type: "INTERNAL_SERVER_ERROR", Code: ERR40002, Message: "Internal error"},
	ERR40003: {Status: http.StatusNotFound, Type: "DATA_NOT_FOUND", Code: ERR40003, Message: "Data Not Found"},
	ERR40004: {Status: http.StatusConflict, Type: "ERR_PROCESSING", Code: ERR40004, Message: "Duplicate process"},
	ERR40005: {Status: http.StatusBadRequest, Type: "ERR_DUPLICATE", Code: ERR40005, Message: "Duplicate Data"},
}
