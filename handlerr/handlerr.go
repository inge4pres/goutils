package hendlerr

import (
	"log"
	"os"
)

type Herr struct {
	LogFile, Level, Descr string
	Err            error
}

func IsErr(err error) error {
	    return err
}

func (e *Herr) HandlErr (level, descr string, err error) {
	e.Err = err
        e.Level = level
        e.LogErr()
}

func (e *Herr) LogErr() {
	l, f := initLog(e.LogFile, e.Level)
	defer f.Close()

	if e.Err != nil {
		l.Println(e.Descr)
		l.Printf("Error: %T - %s", e.Err, e.Err)
	}
}

func initLog(path, level string) (l *log.Logger, file *os.File) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if file, err := os.Create(path); err == nil {
			l = log.New(file, level, log.Lshortfile)
		} else {
			logger := log.New(os.Stdout, "FATAL", log.Llongfile)
			logger.Fatalf("Cannot init log file %s\n%T\n%s", path, err, err)
		}
	return
	}

	file, err := os.OpenFile(path, os.O_RDWR, 0755)

	if err != nil {
		logger := log.New(os.Stdout, "FATAL", log.Llongfile)
		logger.Fatalf("Cannot init log file %s\n%T\n%s", path, err, err)
	}

	l = log.New(file, level, log.Lshortfile)
	return
}
