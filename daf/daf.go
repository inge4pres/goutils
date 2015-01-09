package daf

import (
	"fmt"
	ex "inge4pres/goutils/handlerr"
	sql "database/sql"
	import _ "github.com/go-sql-driver/mysql"
)

const (
	DRVM = "mysql"
)

var e = ex.HanldErr{LogFile: "./databese.err"}

type Daf struct {
	Dbname, Host, Pwd, Port, Protocol, User string
}

type MySqlDaf interface {
	Db 		Daf
	Conn	sql.Open(DRVM, Db.User+":"Db.Pwd+"@"+Db.Protocol+"("+Db.Host+":"+Db.Port+")"+"/"+Db.Dbname) 
}
