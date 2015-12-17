package utils

import (
	"fmt"
	"log"
)

func Log(packageName string, stringFormat string, stringArgs ...interface{}) {
	message := fmt.Sprintf(stringFormat, stringArgs...)
	message = fmt.Sprintf("[%s] %s", packageName, message)
	log.Println(message)
}
