package main

import (
	"fmt"
	"github.com/madhutomy/golang-logging-example/common"
	pkg2 "github.com/madhutomy/golang-logging-example/package2"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var logger *logrus.Logger
var goRoutineMap map[string]string

/**
Set up logrus logger and there-by establish the settings for all packages right away.
*/
func initLogger(logLevel string, formatter string) {
	level, _ := logrus.ParseLevel(logLevel)
	logrus.SetLevel(level)
	logrus.SetFormatter(common.SetFormatter(formatter))
	// To include the calling method as a field
	logrus.SetReportCaller(true)
}

func createSingleLogger(logLevel string, formatter string) *logrus.Logger {
	logger := logrus.New()
	level, _ := logrus.ParseLevel(logLevel)
	logger.SetLevel(level)
	logger.SetFormatter(common.SetFormatter(formatter))
	logrus.SetReportCaller(true)
	return logger
}

func initialize(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
	goRoutineMap = common.InitGoRoutineMaps()
	/*
		Second Approach : Create one instance of a logger, update its log level,
			and then use this logger for all the packages within your project.
			Similar to first approach More manageable and comprehensive
	*/
	logger = createSingleLogger("info", "text")
	testLogger()
	logger.Infof("... Logger Initialization completed with LEVEL=info  ....")
}

func testLogger() {
	fmt.Println("")
	fmt.Println("........ Testing Logger Started .......")
	job2 := pkg2.Job2{Logger: logger, GoRoutineMap: goRoutineMap}
	job2.StartJob2()
	logger.Infof("Info logging...")
	logger.Error("Error logging !!!!")
	logger.Debug("Debug logging...")
	logger.Trace("Trace Logging")
	fmt.Println("........ Testing Logger Completed .......")
	go PeerStatus(logger, goRoutineMap["PEER_STATUS"])
	go PeerSync(logger, goRoutineMap["PEER_SYNC"])
}

func changeLogLevelsAtRuntime(w http.ResponseWriter, req *http.Request) {
	fmt.Println("GET params were:", req.URL.Query())
	param1 := req.URL.Query().Get("logLevel")
	common.ChangeLogLevel(logger, param1)
	testLogger()
}

func main() {
	http.HandleFunc("/init", initialize)
	http.HandleFunc("/change", changeLogLevelsAtRuntime)

	http.ListenAndServe(":8090", nil)

	/*
		3 Approach : You might want to initiate separate loggers with separate logging levels since there are cases
		where you would like to see more logs from one package but avoid seeing more logs from other packages.
	*/

	//createSingleLogger("info", "text")
	//logrus.Println("")
	//logrus.Println(" THIRD APPROACH")
	//job3 := pkg3.NewLogger("debug", "text")
	//job3.StartJob3()
	//
	//logrus.Infof("Approach 3 Info logging...")
	//logrus.Error("Approach 3 Error logging !!!!")
	//logrus.Debug("Approach 3 Debug logging...")
}

func PeerStatus(logger *logrus.Logger, goroutineId string) {
	logEntry := common.CreateLogEntryForConcurrent(logger, goroutineId)

	for _ = range time.Tick(time.Duration(10) * time.Second) {
		logEntry.Errorln("Error")
		logEntry.Traceln("Trace")
		logEntry.Infoln("Info")
		logEntry.Debugln("Debug")
	}

}

func PeerSync(logger *logrus.Logger, goroutineId string) {
	logEntry := common.CreateLogEntryForConcurrent(logger, goroutineId)
	for _ = range time.Tick(time.Duration(20) * time.Second) {
		logEntry.Error("Error")
		logEntry.Trace("Trace")
		logEntry.Info("Info")
		logEntry.Debug("Debug")
		logEntry.WithError(&pkg2.MyError{}).Error("PeerSync failed <> ----!!!")
	}
}
