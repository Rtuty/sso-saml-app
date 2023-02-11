package tools

import (
	"os"
	"strings"
)

// FileExists проверка существует ли файл
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Contains проверяет: входит ли строка в массив строк?
func Contains(arr []string, str string, caseInsensitive bool) bool {
	lowerStr := strings.ToLower(str)
	for _, v := range arr {
		if (v == str) || (caseInsensitive && (strings.ToLower(v) == lowerStr)) {
			return true
		}
	}
	return false
}
