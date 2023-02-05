package log

import "log"

func Info(v ...interface{}) {
	log.Println(v...)
}
