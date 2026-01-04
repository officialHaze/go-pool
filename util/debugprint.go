package util

import (
	"fmt"
	"log"
	"os"
)

type DebugPrint struct {
	msg string
}

func DebugPrinter(msg string) *DebugPrint {
	return &DebugPrint{
		msg: msg,
	}
}

func (dp *DebugPrint) Logln() {
	if os.Getenv("DEBUG") == "1" {
		log.Println(dp.msg)
	}
}

func (dp *DebugPrint) Logf(params ...any) {
	if os.Getenv("DEBUG") == "1" {
		log.Printf(dp.msg, params...)
	}
}

func (dp *DebugPrint) Println() {
	if os.Getenv("DEBUG") == "1" {
		fmt.Println(dp.msg)
	}
}

func (dp *DebugPrint) Printf(params ...any) {
	if os.Getenv("DEBUG") == "1" {
		fmt.Printf(dp.msg, params...)
	}
}
