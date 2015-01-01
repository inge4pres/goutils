package hendlerr

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	INFO  = 0
	WARN  = 1
	DEBUG = 2
	FATAL = 3
)

type Herr struct {
	LogFile, Descr string
	Level          int
	Err            error
}

func (e *Herr) LogErr() {
	l, f := initLog(e.LogFile, e.Level)
	defer f.Close()

	select {

	case e.Level = INFO:
		l.Println("[INFO] " + e.Descr)
		l.Printf("Error :%T - %s", err, err)

	case e.Level = WARN:
		l.Println("[WARN] " + e.Descr)
		l.Printf("Error :%T - %s", err, err)

	case e.Level = DEBUG:

		l.Println("[DEBUG] " + e.Descr)
		l.Printf("Error :%T - %s", err, err)

	case e.Level = INFO:
		l.Println("[FATAL] " + e.Descr)
		l.Fataltf("Error :%T - %s", err, err)
	}

	if e.Err != nil {
		l.Printf("Error: %T - %s", e.Err, e.Err)
	}
}

func initLog(path string, flag int) (l *log.Logger, file *os.File) {

	offset = 0
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if file, err := os.Create(path); err == nil {
			l = log.New(file, nil, flag)
			return
		} else {
			logger := log.New(os.Stdout, nil, log.Llongfile)
			logger.Fatalf("Cannot init log file %s\n%T\n%s", path, err, err)
		}
	}

	file, err := os.OpenFile(path, os.O_RDWR, 0755)

	if err != nil {
		logger := log.New(os.Stdout, nil, log.Llongfile)
		logger.Fatalf("Cannot init log file %s\n%T\n%s", path, err, err)
	}

	l = log.New(file, nil, Lshortfile)
	return
}
