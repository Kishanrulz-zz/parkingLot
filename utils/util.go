package utils

import (
	"os"
)



func ReadFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	return file, err
}


