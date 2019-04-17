package logger

import (
	"log"
	"os"
)

// var (
// 	errorLogger *log.Logger
// 	infoLogger  *log.Logger
// )

var errorLogger = log.New(os.Stdout,
	"ERROR: ",
	log.Ldate|log.Ltime|log.Lshortfile,
)

var infoLogger = log.New(os.Stdout,
	"INFO: ",
	log.Ldate|log.Ltime|log.Lshortfile,
)

// func init() {
// errorF, err := os.OpenFile("error_log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// if err != nil {
// 	log.Fatal("Failed to open log file", err)
// }

// infoF, err := os.OpenFile("info_log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// if err != nil {
// 	log.Fatal("Failed to open log file", err)
// }

// multiError := io.MultiWriter(errorF, os.Stdout)
// multiInfo := io.MultiWriter(infoF, os.Stdout)

// errorLogger = log.New(os.Stdout,
// 	"ERROR: ",
// 	log.Ldate|log.Ltime|log.Lshortfile,
// )

// infoLogger = log.New(os.Stdout,
// 	"INFO: ",
// 	log.Ldate|log.Ltime|log.Lshortfile,
// )
// }

// Error logs passed message
func Error(m string) {
	errorLogger.Println(m)
}

// Info logs passed message
func Info(m string) {
	infoLogger.Println(m)
}
