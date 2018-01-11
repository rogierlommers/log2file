package log2file

import (
	"fmt"
	"os"
	"time"
)

var logFile *os.File

// ActivateForFile sends logging to the filename provided
// if the file already exists, the content will be truncated
func ActivateForFile(targetFile string) error {
	return openFileHandler(targetFile)
}

// Activate sends logging to a file in the current directory with format "2006-01-02-T15-04-05.log"
// if the file already exists, the content will be truncated
func Activate() error {
	targetFile := fmt.Sprintf("%s.log", time.Now().Format("2006-01-02-T15-04-05"))
	return openFileHandler(targetFile)
}

func openFileHandler(targetFilename string) (err error) {
	logFile, err = os.OpenFile(targetFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("logfile creation error: %s", err.Error())
	}

	return nil
}

func Write(message interface{}) {
	m := fmt.Sprintf("%s\n", message.(string))
	logFile.WriteString(m)
}
