package main

import (
	"context"
	"fmt"
	"github.com/madhutomy/golang-logging-example/common"
	pkg2 "github.com/madhutomy/golang-logging-example/package2"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var logger logrus.Logger

//REST end points to initialize the logger and app
func initialize(w http.ResponseWriter, req *http.Request) {
	logger := common.GetLogger()
	logger.Println("Logger Initialized...")
	testLogger()
}

func testLogger() {
	fmt.Println("")
	fmt.Println("........ Testing Logger Started .......")
	pkg2.StartJob()
	logger.Info("Info logging...")
	logger.Error("Error logging !!!!")
	logger.Debug("Debug logging...")
	logger.Trace("Trace Logging")
	fmt.Println("........ Testing Logger Completed .......")
	// Example for a go-routine
	go PeerStatus()
}

// REST endpoint to change the log level
func changeLogLevelsAtRuntime(w http.ResponseWriter, req *http.Request) {
	param := req.URL.Query().Get("logLevel")
	fmt.Println("GET params were:", param)
	logger := common.ModifyLogLevel(param)
	testLogger()
	logger.Infof("... Log Level changed with LEVEL=   ....", logger.GetLevel())
}

func main() {
	http.HandleFunc("/init", initialize)
	http.HandleFunc("/change", changeLogLevelsAtRuntime)
	http.ListenAndServe(":8090", nil)
}

func PeerStatus() {
	ctx := context.Background()
	logger := common.GetLoggerWithContext(common.CtxWithLoggerID(ctx, 1005))
	for now := range time.Tick(time.Duration(10) * time.Second) {
		logger.Info("..... Started the PeerStatus Timer .....:", now)
		logger.Errorln("PeerStatus - Error")
		logger.Traceln("PeerStatus - Trace")
		logger.Infoln("PeerStatus - Info")
		logger.Debugln("PeerStatus - Debug")
	}
}


