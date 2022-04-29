package utils

func InArrar(key string, arr []string) bool {

	for _, v := range arr {
		if key == v {
			return true
		}
	}

	return false
}
