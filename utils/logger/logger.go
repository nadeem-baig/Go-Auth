package logger

import "log"

// Fatal logs a fatal error and exits the program if any argument is non-nil
func Fatal(args ...interface{}) {
	for _, arg := range args {
		if err, ok := arg.(error); ok && err != nil {
			log.Fatal(args...)
		}
	}
}

func Println(msg string)  {
    log.Println("Starting the Go HTTP server...")
	
}

