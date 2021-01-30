package package1

import (
	"context"
	"github.com/madhutomy/golang-logging-example/common"
	log "github.com/sirupsen/logrus"
	"time"
)

type MyError struct {}

func (myErr MyError) Error() string {
	return "boom !!!"
}

func getError() (string, error) {
	return "", &MyError{}
}

func StartJob(){
	logger := common.GetLogger()
	// I created a custom error to test with errors
	_,err := getError()
	logger.Debug("<Debug> Done with Job 2")
	logger.Info("<Info> Done with Job 2")
	logger.Error("<Error> Done with Job 2")
	logger.Trace("<Trace> Done with Job 2")
	go MongoSync()
	go PeerSync()
	// Provide context info when error occurs
	logger1 := logger.WithFields(log.Fields{
		"user_id": 2384172,
		"transaction_id": 9823183,
		"amount": 5000,
		"currency": "USD",
	})
	logger1.Info("Start transaction")
	logger1.WithError(err).Error("Transaction failed <> ----!!!")
}

func MongoSync() {
	ctx := context.Background()
	logger := common.GetLoggerWithContext(common.CtxWithLoggerID(ctx, 3005))
	logger.Info("..... Started the MongoSync Timer .....:")
	logger.Errorln("Error -- MongoSync --")
	logger.Traceln("Trace -- MongoSync --")
	logger.Infoln("Info -- MongoSync --")
	logger.Debugln("Debug -- MongoSync --")

}

func PeerSync() {
	ctx := context.Background()
	logger := common.GetLoggerWithContext(common.CtxWithLoggerID(ctx, 2005))
	for now := range time.Tick(time.Duration(20) * time.Second) {
		logger.Info("..... Started the PeerSync Timer .....:", now)
		logger.Error("Error ## ## PeerSync  ## ##")
		logger.Trace("Trace ## ## PeerSync  ## ##")
		logger.Info("Info ## ## PeerSync  ## ##")
		logger.Debug("Debug ## ## PeerSync  ## ##")
		logger.WithError(&MyError{}).Error("PeerSync failed <> ----!!!")
	}
}
