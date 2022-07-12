package log

import (
	"fmt"
	"github.com/fatedier/beego/logs"
)

var Log *logs.BeeLogger

func init(){
	Log = logs.NewLogger(200)
	Log.EnableFuncCallDepth(true)
	Log.SetLogFuncCallDepth(Log.GetLogFuncCallDepth()+1)
}

func InitLog(logWay string,logfile string, logLevel string, maxdays int64, disableLogColor bool){
	SetLogFile(logWay,logfile,maxdays,disableLogColor)
	SetLogLevel(logLevel)
}

func SetLogFile(logWay string,logFile string,maxdays int64,disableLogColor bool){
	if logWay == "console" {
		params := ""
		if disableLogColor {
			params = fmt.Sprintf(`{"color":false}`)
		}
		Log.SetLogger("console",params)
	}else{
		params := fmt.Sprintf(`{"filename":"%s","maxdays":"%d"}`,logFile,maxdays)
		Log.SetLogger("file",params)
	}
}


func SetLogLevel(logLevel string){
	level := 4
	switch logLevel {
	case "error":
		level=3
	case "warn":
		level=4
	case "info":
		level=6
	case "debug":
		level=7
	case "trace":
		level=8
	default:
		level=4

	}
	Log.SetLevel(level)
}

