package log

import (
	"os"
	"log"
	"path/filepath"
)

const LogDir = "logs"

func Init(name string) {
	filename := filepath.Join(LogDir, name);
	os.MkdirAll(LogDir, os.ModePerm)
	logFile, err := os.OpenFile(filename, os.O_CREATE | os.O_RDWR | os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Println("cannot create log file:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("success create log file:", logFile.Name())
}
