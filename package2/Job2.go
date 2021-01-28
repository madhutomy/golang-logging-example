package package1

import (
	"github.com/madhutomy/golang-logging-example/common"
	log "github.com/sirupsen/logrus"
	"time"
)

type Job2 struct {
	Logger *log.Logger
	GoRoutineMap map[string]string
}

type MyError struct {}

func (myErr MyError) Error() string {
	return "boom !!!"
}

func sayHello() (string, error) {
	return "", &MyError{}
}

func (job Job2) StartJob2(){
	_,err := sayHello()
	job.Logger.Debug("<Debug> Done with Job 2")
	job.Logger.Info("<Info> Done with Job 2")
	job.Logger.Error("<Error> Done with Job 2")
	job.Logger.Trace("<Trace> Done with Job 2")
	go job.MongoSync()
	// Provide context info when error occurs
	logger1 := job.Logger.WithFields(log.Fields{
		"user_id": 2384172,
		"transaction_id": 9823183,
		"amount": 5000,
		"currency": "USD",
	})
	logger1.Info("Start transaction")
	logger1.WithError(err).Error("Transaction failed <> ----!!!")
}

func (job Job2) MongoSync() {
	logEntry := common.CreateLogEntryForConcurrent(job.Logger, job.GoRoutineMap["MONGO_SYNC"])
	for _ = range time.Tick(time.Duration(15) * time.Second) {
		logEntry.Errorln("Error")
		logEntry.Traceln("Trace")
		logEntry.Infoln("Info")
		logEntry.Debugln("Debug")
	}

}
