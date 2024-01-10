package main

import "github.com/alecthomas/kingpin"

var (
	app = kingpin.New("hc", "health checker")
	runCmd = app.Command("run", "run service")
	stopOnFail = app.Flag("stopIfFailed","set to stop ping process if one of url respond with fail").Bool()
	port = app.Arg("port","port to start listening").Default(":8080").String()
	maxTimeout = app.Arg("timeout","set maximum timeout for requests").Default("5").Int()
)