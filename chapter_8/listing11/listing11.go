// this sample program demonstrates how to create customized loggers
package main

import (
	"io"
	"io/iotil"
	"log"
	"os"
)

var (
	Trace   *log.Logger // just about anything
	Info    *log.Logger // important information
	Warning *log.Logger // be concerned
	Error   *log.Logger // critical problem
)

func init() {
	file, err := os.OpenFile("errors.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(iotil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Trace.Println("I have something standart to say")
	Info.Println("Special information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
