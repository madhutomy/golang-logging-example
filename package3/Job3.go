package package1

import (
	"github.com/madhutomy/golang-logging-example/common"
	log "github.com/sirupsen/logrus"
)

type Job3 struct {
	logger *log.Logger
}

func NewLogger(logLevel string, formatter string) Job3 {
	logger := log.New()
	level,_ := log.ParseLevel(logLevel)
	logger.SetLevel(level)
	log.SetFormatter(common.SetFormatter(formatter))
	return Job3{logger: logger}
}

type MyError struct {}

func (myErr MyError) Error() string {
	return "boom !!!"
}

func sayHello() (string, error) {
	return "", &MyError{}
}

func (job Job3) StartJob3(){
	_,err := sayHello()
	job.logger.Debug("Done with Job 3")
	job.logger.Info("Done with Job 3")
	// Provide context info when error occurs
	logger1 := job.logger.WithFields(log.Fields{
		"user_id": 2384172,
		"transaction_id": 9823183,
		"amount": 5000,
		"currency": "USD",
	})
	logger1.Info("Start transaction")
	logger1.WithError(err).Error("Transaction failed <> ----!!!")
}
