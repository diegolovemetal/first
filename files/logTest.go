package main

import (
	"fmt"
	"log"
	"os"
)

//log.Print("this is log print test", "\n")
	//log.Printf("this is %s", "log printf test\n")
	//log.Printf("this is log println test")
	//defer func() {
	//	fmt.Println("first defer")
	//}()
	//log.Fatal("this is log fatal test","\n")
	//log.Fatalf("this is %s","log fatalf test")
	//log.Fatalln("this is fatalln test\n")
	//defer func() {
	//	fmt.Println("second defer")
	//}()
	//func main() {
	//	defer func() {
	//		fmt.Println("first defer")
	//		if err := recover(); err != nil {
	//			fmt.Println(err)
	//		}
	//	}()
	//	log.Panic("this is log panic")
	//	defer func() {
	//		fmt.Println("second defer")
	//	}()
	//}
func Debug(logName string) {
	logFile, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("create ./test.log err : %v\n", err)
	}
	if logFile != nil {
		defer func(file *os.File) { file.Close() }(logFile)
	}

	debugLog := log.New(logFile, "[Debug]", log.Ldate)

	debugLog.SetPrefix("[Debug]")
	debugLog.SetFlags(log.Lshortfile)
	debugLog.Println("this is Debug log")
}
func Waring(logName string) {
	logFile, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("create ./test.log err : %v\n", err)
	}
	if logFile != nil {
		defer func(file *os.File) { file.Close() }(logFile)
	}

	debugLog := log.New(logFile, "[Waring]", log.Ldate)

	debugLog.SetPrefix("[Waring]")
	debugLog.SetFlags(log.Lshortfile)
	debugLog.Println("this is Waring log")
}

func main() {
	logName := "./test.log"
	Debug(logName)
	Waring(logName)
}
