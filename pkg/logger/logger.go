package logger

import (
	"io"
	"log"
	"os"
)

var (
	errorLogger *log.Logger
	infoLogger  *log.Logger
)

func init() {
	errorF, err := os.OpenFile("error_log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file", err)
	}

	infoF, err := os.OpenFile("info_log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file", err)
	}

	multiError := io.MultiWriter(errorF, os.Stdout)
	multiInfo := io.MultiWriter(infoF, os.Stdout)

	errorLogger = log.New(multiError,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	infoLogger = log.New(multiInfo,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
}

// Error logs passed message
func Error(m string) {
	errorLogger.Println(m)
}

// Info logs passed message
func Info(m string) {
	infoLogger.Println(m)
}
