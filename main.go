package main

import (
	"io"
	"log"
	"os"
	"gopool/settings"
	"gopool/util"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	// Generate my settings
	if err := settings.Generate(); err != nil {
		log.Fatalln(err)
	}

	// load the env file
	if err := godotenv.Load(settings.MySettings.ENV_FILE_NAME); err != nil {
		log.Fatalln(err)
	}

	// Setup lumberjack and logger
	logfile := &lumberjack.Logger{
		Filename:   "/var/log/gopool.log",
		MaxSize:    100, // mb
		MaxBackups: 5,
		MaxAge:     28, // days
		Compress:   true,
	}

	// Log on both console and logfile when in dev mode
	if util.InDevMode() {
		multiwriter := io.MultiWriter(os.Stdout, logfile)
		log.SetOutput(multiwriter)
	} else {
		log.SetOutput(logfile)
	}
}

func main() {
	log.Printf(" ****** PROJECT Running In %s Environment ******", strings.ToUpper(os.Getenv("ENV")))
}
