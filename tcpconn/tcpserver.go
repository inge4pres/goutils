package tcpconn

import (
	"fmt"
	ex "inge4pres/goutils/handlerr"
	"net"
	"os"
	json "encoding/json"
	"errors"
)
/*
Declared global logfile
*/
var e = ex.Herr{LogFile: "./tcpconnection.err"}

/*
Main function called from other referenced packages and main; it opens the socket and listen for clients.
Should be extended in "handleConnection" to have a full protocol implementation
*/
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
			e.HandlErr("WARN ", "Could not start the server", err)
			return err
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn *net.TCPConn) {
	comm, err := startConn(conn)
	if err != nil {
		e.HandlErr("WARN ","Error initiating the connection!", err)
	}
	err = initHandShake(conn, comm)
	if err != nil {
		e.HandlErr("WARN ", "Phase "+string(comm.Phase)+" went WRONG!", err)
	}
	err, resp := execCommand(comm)
	if err != nil {
		e.HandlErr("WARN ", "Command execution failed...", err)
	}
	fmt.Println("RESPONSE GENERATED: "+string(resp))
}

func startConn(conn *net.TCPConn) (comm *TCPCommand, ex error) {
	var message = make(chan []byte, WRKR_COUNT)
	var erCh = make(chan error)
	go func(){
		read := make([]byte, MAX_COMM_SIZE)
		b, err := conn.Read(read)
		message<-read[0:b]
		erCh<-err
	}()
	buf := <-message
	err := <-erCh
	if err != nil {
		e.HandlErr("WARN ", "This request generated an error!\n"+string(buf), err)
	}
	comm = &TCPCommand{}
	err = json.Unmarshal(buf, comm)
	if err != nil {
		e.HandlErr("WARN ", "Could not understand the received command!", err)
	}
/*	Implements only the initial handshake
	Managing the rest of portocol after handshake must be done in specific way
*/	return comm, err
}

func initHandShake(conn *net.TCPConn, comm *TCPCommand) error {
	if comm.Phase == PHASE_CONN {
		conn.Write([]byte(RESP_OK))
		return nil
	}
	return errors.New("Could not start handshake with client, protocol phase mismatch")
}

func execCommand(comm *TCPCommand) (err error, resp []byte) {
	//TODO choose waht to do with the command received
	return nil, []byte(RESP_OK)
}
