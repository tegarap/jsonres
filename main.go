package util

const (
	statusSuccess = "success"
	statusFail    = "fail"
)

func ResponseSuccess(message string, data interface{}) map[string]interface{} {
	if data != nil {
		return map[string]interface{}{
			"status": statusSuccess,
			"message": message,
			"data": data,
		}
	}
	return map[string]interface{}{
		"status": statusSuccess,
		"message": message,
	}
}
func ResponseFail(message string, data interface{}) map[string]interface{} {
	if data != nil {
		return map[string]interface{}{
			"status": statusFail,
			"message": message,
			"data": data,
		}
	}
	return map[string]interface{}{
		"status": statusFail,
		"message": message,
	}
}
