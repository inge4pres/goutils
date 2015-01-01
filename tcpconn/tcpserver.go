package  tcpconn

import (
	"fmt"
	"net"
	"os"
)

func errorHandler(e error) bool {
	res := true
	if(e != nil){
		res = false
	}
	return res
}

func TcpServerListener(conntype, host, port string) error {
	addr, _ := net.ResolveTCPAddr(conntype, host+":"+port)
	l, err := net.ListenTCP(conntype, addr)
	defer l.Close()
	if(errorHandler(err)){
		fmt.Fprintf(os.Stdout,"Accepting connection on %s://%s:%s\n", conntype, host, port )
	} else{
		return err
	}
	for {
		conn, err := l.AcceptTCP()
		if (!errorHandler(err)){
			fmt.Fprintf(os.Stdout,"Error occured during connection: %s\n", err)
			return err
		}
	go handleConnection(conn)
	}
}


func handleConnection(conn *net.TCPConn) (ex error){
	var message = make(chan []byte, WRKR_COUNT)
	go func(){
		read := make([]byte, MAX_COMM_SIZE)
		_, err := conn.Read(read)
		message <- read
		if(!errorHandler(err)){
			fmt.Fprintf(os.Stdout, "The current request generateed error!\n%s\n%T - %s", string(read), err, err)
			ex = err
			return
		}
	}()
	buf := <- message
	phase, err := handlePhase(conn, buf)
	if(!errorHandler(err)){
		fmt.Fprintf(os.Stdout, "There was an error communicating during phase %d connecting with %s", phase, conn.RemoteAddr().String())
	}
	conn.Close()
	ex = err
	return
}



func handlePhase(conn *net.TCPConn, buf []byte) (phase string, ex error){
	phase = string(buf)
	fmt.Fprintf(os.Stdout, "PHASE_RECV: %s\n", phase)
	return
}
