package api

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	Warning      *log.Logger
	Info         *log.Logger
	Debug        *log.Logger
	Error        *log.Logger
	Gin          *log.Logger
	Stats        *log.Logger
	errorColor   = color.New(color.Bold, color.FgRed).SprintFunc()
	infoColor    = color.New(color.Bold, color.FgWhite).SprintFunc()
	debugColor   = color.New(color.Bold, color.FgGreen).SprintFunc()
	warningColor = color.New(color.Bold, color.FgYellow).SprintFunc()
	versionColor = color.New(color.Bold, color.FgHiCyan).SprintFunc()
	ginColor     = color.New(color.Bold, color.FgBlue).SprintFunc()
	wrt          io.Writer
)

var route string

func InitLog(debugEnabled bool, desiredLogRoute, version string) {
	route = desiredLogRoute
	f, err := os.OpenFile(route, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	wrt = io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)

	Info = log.New(wrt, infoColor("\n[INFO] "), log.Ldate|log.Ltime|log.Lshortfile)
	Info.SetOutput(wrt)

	Warning = log.New(wrt, warningColor("\n[WARNING] "), log.Ldate|log.Ltime|log.Lshortfile)
	Warning.SetOutput(wrt)

	Debug = log.New(wrt, debugColor("\n[DEBUG] "), log.Ldate|log.Ltime|log.Lshortfile)
	Debug.SetOutput(wrt)
	if !debugEnabled {
		var buff bytes.Buffer
		Debug.SetOutput(&buff)
	}

	Error = log.New(wrt, errorColor("\n[ERROR] "), log.Ldate|log.Ltime|log.Lshortfile)
	Error.SetOutput(wrt)

	Gin = log.New(wrt, ginColor("\n[GIN] "), log.Ldate|log.Ltime)
	Gin.SetOutput(wrt)

	Stats = log.New(wrt, debugColor("\n[STATS] "), log.Ldate|log.Ltime|log.Lshortfile)
	Stats.SetOutput(wrt)
	if !debugEnabled {
		var buff bytes.Buffer
		Stats.SetOutput(&buff)
	}

	PrintVersion(version)
}

func PrintVersion(version string) {
	os.Remove("SDK.ver")
	f, err := os.OpenFile("SDK.ver", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	versionLog := log.New(os.Stdout, versionColor("\n[VERSION] "), 0)
	versionLog.SetOutput(os.Stdout)
	versionLog.Println("CERBERUS", version, versionColor("[VERSION]"))

	vLog := log.New(f, "", 0)
	vLog.SetOutput(f)

	vLog.Println("CERBERUS VERSION:", version)
}

func EnableDebug(activated bool) {
	f, err := os.OpenFile(route, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	wrt = io.MultiWriter(os.Stdout, f)
	var buff bytes.Buffer
	if activated {
		Debug.SetOutput(wrt)
	} else {
		Debug.SetOutput(&buff)
	}
}