package app

import (
	"log"
	"os"

	"github.com/sriram-yeluri/nxrm3/util"
)

var NXRMConfig util.Config

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func writeLogsToFile() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "[INFO] : ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "[WARNING] : ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "[ERROR] : ", log.Ldate|log.Ltime|log.Lshortfile)
}

func writeLogsToStdOut() {
	InfoLogger = log.New(os.Stdout, "[INFO] : ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "[WARNING] : ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "[ERROR] : ", log.Ldate|log.Ltime|log.Lshortfile)
}
func init() {

	writeLogsToFile()
	//writeLogsToStdOut()

	NXRMConfig = util.LoadConfig()
	InfoLogger.Println(NXRMConfig)
}
