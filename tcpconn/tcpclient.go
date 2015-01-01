package tcpconn

import (
	"net"
	json "encoding/json"
	"os"
	"fmt"
)

func (comm *TCPCommand) Connect4() (conn *net.TCPConn, ex error) {
	addr, _ := net.ResolveTCPAddr("tcp4", comm.RHost+":"+comm.RPort)
	conn, ex = net.DialTCP("tcp", nil, addr)
	if (ex != nil){
		fmt.Fprintf(os.Stdout, "Error during the connection to server %si:%s", comm.RHost, comm.RPort)
		fmt.Fprintf(os.Stdout, "Error is %T\n%s", ex, ex)
		os.Exit(1)
	}
	return
}

func (comm *TCPCommand) SendCommand(conn *net.TCPConn) {
	conn.Write([]byte(START_COMM))
	buf := make([]byte, MAX_BUFF_SIZE)
	_, err := conn.Read(buf)
	if (err != nil) {
		fmt.Fprintf(os.Stdout, "Error reading from buffer of TCP Connection\n%T\n%s\n", err, err)
		os.Exit(1)
	}
	resp := string(buf)
	if (resp == RESP_OK) {
		buf = make([]byte, MAX_BUFF_SIZE)
		jenc, _ := json.Marshal(comm)
		conn.Write([]byte(jenc))
	}
}

