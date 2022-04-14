package nelog

import (
	"fmt"
	"newe-serve/common/setting"
	"newe-serve/pkg/file"
	"sync"
	"time"

	"github.com/aiwuTech/fileLogger"
)

var (
	LOGTIME  string
	FILENAME string
	logFile  *fileLogger.FileLogger
	TRACE    *fileLogger.FileLogger
	INFO     *fileLogger.FileLogger
	WARN     *fileLogger.FileLogger
	ERROR    *fileLogger.FileLogger
	CRON     *fileLogger.FileLogger
	WG       *sync.WaitGroup
)

func SetUp() {
	LOGTIME = time.Now().Format("20060102")
	FILENAME := fmt.Sprintf("%s.%s", time.Now().Format("20060102"), "log")
	file.IsNotExistMkDir(setting.SYS.RuntimeRootPath)

	tracepath := setting.SYS.RuntimeRootPath + "/trace/" + time.Now().Format("200601")
	if err := file.IsNotExistMkDir(tracepath); err != nil {
		fmt.Println("日志创建失败")
	}
	TRACE = fileLogger.NewSizeLogger(tracepath, FILENAME, "-", 10, 2, fileLogger.MB, 300, 5000)

	infopath := setting.SYS.RuntimeRootPath + "/info/" + time.Now().Format("200601")
	if err := file.IsNotExistMkDir(infopath); err != nil {
		fmt.Println("日志创建失败")
	}
	INFO = fileLogger.NewSizeLogger(infopath, FILENAME, "-", 10, 2, fileLogger.MB, 300, 5000)

	warnpath := setting.SYS.RuntimeRootPath + "/warn/" + time.Now().Format("200601")
	if err := file.IsNotExistMkDir(warnpath); err != nil {
		fmt.Println("日志创建失败")
	}
	WARN = fileLogger.NewSizeLogger(warnpath, FILENAME, "-", 10, 2, fileLogger.MB, 300, 5000)

	errorpath := setting.SYS.RuntimeRootPath + "/error/" + time.Now().Format("200601")
	if err := file.IsNotExistMkDir(errorpath); err != nil {
		fmt.Println("日志创建失败")
	}
	ERROR = fileLogger.NewSizeLogger(errorpath, FILENAME, "-", 10, 2, fileLogger.MB, 300, 5000)

	cronpath := setting.SYS.RuntimeRootPath + "/cron/" + time.Now().Format("200601")
	if err := file.IsNotExistMkDir(cronpath); err != nil {
		fmt.Println("日志创建失败")
	}
	CRON = fileLogger.NewSizeLogger(cronpath, FILENAME, "-", 10, 2, fileLogger.MB, 300, 5000)

	TRACE.SetPrefix("[TRACE] ")
	INFO.SetPrefix("[INFO] ")
	WARN.SetPrefix("[WARN] ")
	ERROR.SetPrefix("[ERROR] ")
	CRON.SetPrefix("[CRON] ")
	WG = new(sync.WaitGroup)
	defer func() {
		fmt.Println("日志OVER")
	}()

}

func Trace(v ...interface{}) {
	WG.Wait()
	if LOGTIME != time.Now().Format("20060102") {
		WG.Add(1)
		setnew()
	}
	TRACE.Println(v...)
	return
}

func Info(v ...interface{}) {
	WG.Wait()
	if LOGTIME != time.Now().Format("20060102") {
		WG.Add(1)
		setnew()
	}
	INFO.Println(v...)
	return
}

func Warn(v ...interface{}) {
	WG.Wait()
	if LOGTIME != time.Now().Format("20060102") {
		WG.Add(1)
		setnew()
	}
	WARN.Println(v...)
	return
}

func Error(v ...interface{}) {
	WG.Wait()
	if LOGTIME != time.Now().Format("20060102") {
		WG.Add(1)
		setnew()
	}
	ERROR.Println(v...)
	return
}

func Cron(v ...interface{}) {
	WG.Wait()
	if LOGTIME != time.Now().Format("20060102") {
		WG.Add(1)
		setnew()
	}
	CRON.Println(v...)
	return
}

func setnew() {
	TRACE.Close()
	INFO.Close()
	WARN.Close()
	ERROR.Close()
	LOGTIME = time.Now().Format("20060102")
	FILENAME = fmt.Sprintf("%s.%s", time.Now().Format("20060102"), "log")
	tracepath := setting.SYS.RuntimeRootPath + "/trace/" + time.Now().Format("200601")
	if err := file.IsNotExistMkDir(tracepath); err != nil {
		fmt.Println("日志创建失败")
	}
	TRACE = fileLogger.NewSizeLogger(tracepath, FILENAME, "-", 10, 2, fileLogger.MB, 300, 5000)

	infopath := setting.SYS.RuntimeRootPath + "/info/" + time.Now().Format("200601")
	if err := file.IsNotExistMkDir(infopath); err != nil {
		fmt.Println("日志创建失败")
	}
	INFO = fileLogger.NewSizeLogger(infopath, FILENAME, "-", 10, 2, fileLogger.MB, 300, 5000)

	warnpath := setting.SYS.RuntimeRootPath + "/warn/" + time.Now().Format("200601")
	if err := file.IsNotExistMkDir(warnpath); err != nil {
		fmt.Println("日志创建失败")
	}
	WARN = fileLogger.NewSizeLogger(warnpath, FILENAME, "-", 10, 2, fileLogger.MB, 300, 5000)

	errorpath := setting.SYS.RuntimeRootPath + "/error/" + time.Now().Format("200601")
	if err := file.IsNotExistMkDir(errorpath); err != nil {
		fmt.Println("日志创建失败")
	}
	ERROR = fileLogger.NewSizeLogger(errorpath, FILENAME, "-", 10, 2, fileLogger.MB, 300, 5000)

	cronpath := setting.SYS.RuntimeRootPath + "/cron/" + time.Now().Format("200601")
	if err := file.IsNotExistMkDir(cronpath); err != nil {
		fmt.Println("日志创建失败")
	}
	CRON = fileLogger.NewSizeLogger(cronpath, FILENAME, "-", 10, 2, fileLogger.MB, 300, 5000)

	TRACE.SetPrefix("[TRACE] ")
	INFO.SetPrefix("[INFO] ")
	WARN.SetPrefix("[WARN] ")
	ERROR.SetPrefix("[ERROR] ")
	CRON.SetPrefix("[CRON] ")
	WG.Done()
	return
}
