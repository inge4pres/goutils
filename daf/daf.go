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

type MySqlDaf struct {
	DbInfo	Daf
	Conn	*sql.DB
}

func (mdb *MySqlDaf) ConnectTo() (err error) {
	myDaf := mdb.DbInfo
	mdb.Conn, err = sql.Open(DRVM, myDaf.User+":"+myDaf.Pwd+"@"+myDaf.Protocol+"("+myDaf.Host+":"+myDaf.Port+")"+"/"+myDaf.Dbname)
	if err!= nil {
		return err
	}
}
