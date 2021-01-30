# golang-logging-example
This is a sample project that shows the usage of the prominent logging library in golang [logrus](https://github.com/sirupsen/logrus)

### Scenarios covered
1. How to re-use the same logger across various packages
2. How we can add a unique loggerId for go routines.
3. How to change the log level at runtime

### To run the project
`go run main.go`
Then initialize by 
`http://localhost:8090/init`
Note:By default, the log level will be info

### How to change log level at runtime
1. I am exposing a HTTP REST end point to change the log level at run time. 
2. You can use query parameter logLevel to change the log level, possible values :- trace,error,info,debug
`http://localhost:8090/change?logLevel=trace`

### How to re-use the same logger across various packages
Followed a Singleton pattern, where there is a single instance of logger available for the app.

```go
// This is in common/logger.go file
func GetLogger() *logrus.Logger {
	lock.Lock()
	defer lock.Unlock()

	if logger == nil {
		logger = logrus.New()
		logger.SetLevel(logrus.InfoLevel)
		logger.SetFormatter(SetFormatter("text"))
		logger.SetReportCaller(true)
	}

	return logger
}
```
### How to access the logger
`logger := common.GetLogger()`
```go
func StartJob2(){
	logger := common.GetLogger()
	logger.Debug("<Debug> Done with Job 2")
	logger.Info("<Info> Done with Job 2")
	logger.Error("<Error> Done with Job 2")
	logger.Trace("<Trace> Done with Job 2")
}
```
### How we can add a unique loggerId for go routines.
Check the following function in Job.go file
```go
func MongoSync() {
	ctx := context.Background()
	logger := common.GetLoggerWithContext(common.CtxWithLoggerID(ctx, 3005))
	logger.Info("..... Started the MongoSync Timer .....:")
	logger.Errorln("Error -- MongoSync --")
	logger.Traceln("Trace -- MongoSync --")
	logger.Infoln("Info -- MongoSync --")
	logger.Debugln("Debug -- MongoSync --")

}
```