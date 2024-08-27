package errors

import (
	"regexp"
	"strconv"
)

var (
	regErrCode = regexp.MustCompile(`\[errcode:(\d+)\]`)
)

// ExtractErrCode 从 error 消息中提取 uint32 的错误码
func ExtractErrCode(err error) (uint32, bool) {
	matches := regErrCode.FindStringSubmatch(err.Error())
	if len(matches) <= 1 {
		return 0, false
	}

	code := matches[1]
	if c, err := strconv.ParseUint(code, 10, 32); err != nil {
		return uint32(c), true
	}
	return 0, false
}
