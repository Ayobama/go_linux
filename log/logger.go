package log

import (
	"fmt"
	"log"
	"os"
)

//type Linux_logger interface {
//	Info()
//}

func overLogger(log string) {
	//var overLogger log.Logger
	infoLogger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}


//OverLord.
//Creates a new logger and ensure it logs into ./Creation/.logs/linuxlogs.log.
//Format: "golinux:", Date, Time, File, Error.
func Log(logg string) error {

	name := "linuxlogs"
	err := os.MkdirAll("C:/Users/dell/github/golang/src/learning/go_linux/log/.logs/", 0777)
	if err != nil {
		err := overLogger(fmt.Sprintln("LOGGER ERROR: Folder could not be created: ", err))
		return err
	}

	var file *os.File
	file, err = os.OpenFile("./log/.logs/"+name+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	//customLogger.SetOutput(file)
	if err != nil {
		overLogger.Println("LOGGER ERROR: File could not open: ", err))
		return err
	}
	defer file.Close()

	overLogger.Println(logg)
	return nil
}

func preSeed(initVar []string, overLogger *log.Logger) {
	for _, err := range initVar {
		overLogger.Println(err)
	}
}


func Info(logg error, overLogger *log.Logger) error {
	var initVar []string
	name := "info"

	err := os.MkdirAll("C:/Users/dell/github/golang/src/learning/go_linux/log/.logs/", 0777)
	if err != nil {
		overLogger.Printf("(%s)INFO ERROR: Folder could not be created: %v\n", name, err)
		return err
	}

	var file *os.File
	file, err = os.OpenFile("./log/.logs/"+name+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	//customLogger.SetOutput(file)
	if err != nil {
		overLogger.Printf("(%s)INFO ERROR: info.log could not open: %v\n", name, err)
		return err
	}
	defer file.Close()

	infoLogger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger.Println(logg)
	return nil
}

//next create a function to be able to log the contents of the file to os.StdOut

func writetofile(logg error, customLogger *log.Logger) error {
	name := "linuxlogs"
	err := os.MkdirAll("C:/Users/dell/github/golang/src/learning/go_linux/log/.logs/", 0777)
	if err != nil {
		customLogger.Println("LOGGER ERROR: Folder could not be created: ", err)
		return err
	}

	var file *os.File
	file, err = os.OpenFile("./log/.logs/"+name+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	customLogger.SetOutput(file)
	if err != nil {
		customLogger.Println("LOGGER ERROR: File could not open: ", err)
		return err
	}
	defer file.Close()

	customLogger.Println(logg)
	return nil
}
