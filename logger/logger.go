package logger

import (
	"fmt"
	"log"
)

type Logger struct {
	prefix string
}

func INFO() *Logger {
	return &Logger{
		prefix: "[INFO]",
	}
}

func SUCCESS() *Logger {
	return &Logger{
		prefix: "[OK]",
	}
}

func WARN() *Logger {
	return &Logger{
		prefix: "[WARN]",
	}
}

func ERROR() *Logger {
	return &Logger{
		prefix: "[ERROR]",
	}
}

func (l *Logger) Println(a ...any) {
	for _, v := range a {
		log.Printf("%s	%v\n", l.prefix, v)
	}
}

func (l *Logger) Printf(f string, v ...any) {
	o := fmt.Sprintf("%s	%s\n", l.prefix, f)
	log.Printf(o, v...)
}
