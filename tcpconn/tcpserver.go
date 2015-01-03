package tcpconn

import (
	"fmt"
	ex "inge4pres/goutils/handlerr"
	"net"
	"os"
)

var e = ex.Herr{LogFile: "./tcpserver.err"}

func TcpServerListener(conntype, host, port string) error {
	addr, _ := net.ResolveTCPAddr(conntype, host+":"+port)
	l, err := net.ListenTCP(conntype, addr)
	defer l.Close()
	if ex.IsErr(err) == nil {
		fmt.Fprintf(os.Stdout, "Accepting connection on %s://%s:%s\n", conntype, host, port)
	} else {
		e.HandlErr("FATAL", "Could not open socket for listening, exiting...", err)
		return err
	}
	for {
		conn, err := l.AcceptTCP()
		if ex.IsErr != nil {
			e.HandlErr("WARN", "", err)
			return err
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn *net.TCPConn) error {
	var message = make(chan []byte, WRKR_COUNT)
	var erCh = make(chan error)
	go func() {
		read := make([]byte, MAX_COMM_SIZE)
		_, err := conn.Read(read)
		message <- read
		erCh <- err
	}()
	buf := <-message
	err := <-erCh
	if err != nil {
		e.HandlErr("WARN", "This request generateed an error!\n"+string(buf), err)
	}
	phase, err := checkConn(conn, buf)
	if err != nil {
		e.HandlErr("WARN", "There was an error communicating during phase "+string(phase)+" connecting with "+conn.RemoteAddr().String()+"\n", err)
	}
	conn.Close()
	//TODO Managing conncetion
	return err
}

func checkConn(conn *net.TCPConn, buf []byte) (phase string, ex error) {
	phase = string(buf)
	fmt.Fprintf(os.Stdout, "PHASE_RECV: %s\n", phase)
	conn.Write([]byte(RESP_OK))
	return
}
