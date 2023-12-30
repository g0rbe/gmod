package log

import (
	"os"
)

func ExampleLog() {

	Log(LOG_INFO, os.Stdout, "Hello, %s!", "World")
}
