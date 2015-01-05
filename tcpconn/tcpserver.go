package tcpconn

import (
	"fmt"
	ex "inge4pres/goutils/handlerr"
	"net"
	"os"
	json "encoding/json"
	"errors"
)

var e = ex.Herr{LogFile: "./tcpconnection.err"}

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
		if err != nil {
			e.HandlErr("WARN", "", err)
			return err
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn *net.TCPConn) {
	comm, err := StartConn(conn)
//TODO The rest of protocol specification
	if err != nil {
	fmt.Fprintf(os.Stdout, "Error during connection!\n%T\n%s\nClient info:\n%s\n%s", err, err, comm.LAddr, comm.LHost)
		}
	fmt.Fprintf(os.Stdout, "Received command:\nPhase:"+string(comm.Phase)+"\nOS:"+comm.OS+"\nRaddr"+comm.RAddr)
}

func StartConn(conn *net.TCPConn) (comm *TCPCommand, ex error) {
	var message = make(chan []byte, WRKR_COUNT)
	var erCh = make(chan error)
	for {
		read := make([]byte, MAX_COMM_SIZE)
		b, err := conn.Read(read)
		message<-read
		erCh<-err
		if b == 0 {
			break
		}
	}
	buf := <-message
	err := <-erCh
	if err != nil {
		e.HandlErr("WARN", "This request generateed an error!\n"+string(buf), err)
	}
	json.Unmarshal(buf, comm)
	err = startHandShake(conn, comm)
	if err != nil {
		e.HandlErr("WARN", "There was an error communicating during phase "+string(comm.Phase)+" connecting with "+conn.RemoteAddr().String()+"\n", err)
	}
/*	Implements only the initial handshake
	Managing the rest of portocol after handshake must be done in specific way
*/	return comm, err
}

func startHandShake(conn *net.TCPConn, comm *TCPCommand) error {
	if comm.Phase == PHASE_CONN {
		conn.Write([]byte(RESP_OK))
		return nil
	}
	return errors.New("Could not start handshake with client, protocol phase mismatch")
}
