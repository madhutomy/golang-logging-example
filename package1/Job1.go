package package1

import log "github.com/sirupsen/logrus"

type MyError struct {

}
func (myErr MyError) Error() string {
	return "boom !!!"
}

func sayHello() (string, error) {
	return "", &MyError{}
}

func StartJob1(){
	// to be in the main pgm
	// First approach
	//initLogger("info", "json")
	//log.Println("")
	//log.Println(" FIRST APPROACH")
	//package1.StartJob1()
	//logrus.Infof("Approach 1 Info logging...")
	//logrus.Error("Approach 1 Error logging !!!!")
	//logrus.Debug("Approach 1 Debug logging...")
	_,err := sayHello()
	log.Debug("Done with Job 1")
	// Provide context info when error occurs
	logger := log.WithFields(log.Fields{
		"user_id": 2384172,
		"transaction_id": 9823183,
		"amount": 5000,
		"currency": "USD",
	})
	logger.Info("Start transaction")
	logger.WithError(err).Error("Transaction failed <> ----!!!")
}
