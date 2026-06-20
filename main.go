package main

import (
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	// Setup lumberjack and logger
	logfile := &lumberjack.Logger{
		Filename:   "/var/log/gopool.log",
		MaxSize:    50, // mb
		MaxBackups: 5,
		MaxAge:     28, // days
		Compress:   true,
	}

	// Log on both console and logfile
	multiwriter := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(multiwriter)
}

func main() {
}
