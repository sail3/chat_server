package utils

// TO ternary operator
func TO(condition bool, value1 interface{}, value2 interface{}) interface{} {
	if condition {
		return value1
	}
	return value2
}
