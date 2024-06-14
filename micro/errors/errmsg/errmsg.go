package errmsg

import "fmt"

// MissingField 用于 field 漏传递了
func MissingField(field string) string {
	return fmt.Sprintf("%q is required", field)
}

// ZeroField 用于 field 的内容为 0
func ZeroField(field string) string {
	return fmt.Sprintf("%q shouldn't be 0", field)
}

// NilField 用于 field 传递了空值 nil
func NilField(field string) string {
	return fmt.Sprintf("%q shouldn't be nil", field)
}

// EmptyField 用于 field 的内容为空，如空字符串、长度为 0 的数组
func EmptyField(field string) string {
	return fmt.Sprintf("%q shouldn't be empty", field)
}

// OutOfRange 用于 field 的值超出值域范围
func OutOfRange(field string, value interface{}) string {
	return fmt.Sprintf("value %q of %q is out of range", value, field)
}

// InvalidValue 用于 field 使用了非法的值
func InvalidValue(field string, value interface{}) string {
	return fmt.Sprintf("value %q of %q is invalid", value, field)
}

// TimeShouldBeEarlier 用于 field 时间比较，当前 value 应当早于 compareTo
func TimeShouldBeEarlier(field string, value interface{}, compareTo interface{}) string {
	return fmt.Sprintf("time %q of %q should be earlier than %q", value, field, compareTo)
}

// TimeShouldBeEarlierOrEqual 用于 field 时间比较，当前 value 应当早于或等于 compareTo
func TimeShouldBeEarlierOrEqual(field string, value interface{}, compareTo interface{}) string {
	return fmt.Sprintf("time %q of %q should be earlier or equal than %q", value, field, compareTo)
}

// TimeShouldBeLater 用于 field 时间比较，当前 value 应当晚于 compareTo
func TimeShouldBeLater(field string, value interface{}, compareTo interface{}) string {
	return fmt.Sprintf("time %q of %q should be later than %q", value, field, compareTo)
}

// TimeShouldBeLaterOrEqual 用于 field 时间比较当前 value 应当晚于或等于 compareTo
func TimeShouldBeLaterOrEqual(field string, value interface{}, compareTo interface{}) string {
	return fmt.Sprintf("time %q of %q should be later or equal than %q", value, field, compareTo)
}
