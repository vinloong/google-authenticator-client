package authenticator

import (
	"regexp"
	"strings"
)

var COMM_SUBSTRING = []string{`\+`, "=", "/", "\\", "-", " "}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

/**
  替换多个字符串
  TODO 部分字符需要转义 如 + “\+”
*/
func Replace(raw string, cutsets []string, new string) string {
	regStr := strings.Join(cutsets, "|")
	reg := regexp.MustCompile(regStr)
	str := reg.ReplaceAllString(raw, new)
	return str
}

func Trim(raw string, cutsets []string) string {
	return Replace(raw, cutsets, "")
}

func TrimCOMM(raw string) string {
	return Replace(raw, COMM_SUBSTRING, "")
}

func splice(str1, str2 string) (string, error) {
	builder := strings.Builder{}

	_, err := builder.WriteString(str1)
	if err != nil {
		return builder.String(), err
	}
	_, err = builder.WriteString(str2)
	if err != nil {
		return builder.String(), err
	}
	return builder.String(), nil
}
