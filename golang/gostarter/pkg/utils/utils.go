package utils

import (
	"os"
)

func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Log(message string) {
	// Open file for writing
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	// Write message to file
	if _, err := f.Write([]byte(message + "\n")); err != nil {
		panic(err)
	}
	// Close file
	if err := f.Close(); err != nil {
		panic(err)
	}
}
