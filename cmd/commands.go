package main

import "github.com/alecthomas/kingpin"

var (
	app = kingpin.New("app", "health checker")
	runCmd = app.Command("run", "run service")
	stopOnFail = runCmd.Flag("stopIfFailed","set to stop ping process if one of url respond with fail").Bool()
	port = runCmd.Flag("port","port to start listening").Default(":8080").String()
	maxTimeout = runCmd.Flag("timeout","set maximum timeout for requests").Default("5").Int()
)