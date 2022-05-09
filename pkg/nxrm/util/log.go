package util

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	DebugLogger *log.Logger
	ErrorLogger *log.Logger
)

func Info() *log.Logger {
	InfoLogger = log.New(os.Stdout, "[INFO] : ", log.Ldate|log.Ltime|log.Lshortfile)
	return InfoLogger
}

func Debug() *log.Logger {
	DebugLogger = log.New(os.Stdout, "[INFO] : ", log.Ldate|log.Ltime|log.Lshortfile)
	return DebugLogger
}

func Error() *log.Logger {
	ErrorLogger = log.New(os.Stdout, "[ERROR] : ", log.Ldate|log.Ltime|log.Lshortfile)
	return ErrorLogger
}
