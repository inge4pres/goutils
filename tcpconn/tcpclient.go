package tcpconn

import (
	"net"
	json "encoding/json"
	"os"
	"fmt"
	"errors"
	"bytes"
)

//TODO Add IPv6 support
/*
Connect4: starts a IPv4 connection with a remote server, leaves it open for future use 
*/
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
/*
Disconnect: close the previously opened connection
*/
func (comm *TCPCommand) Disconnect(conn *net.TCPConn) {
	conn.Close()
}
/*
PostCommand: send a command on a previously opened connection; command is json-encoded; returns the response in a TcpExData interface
*/
func (comm *TCPCommand) PostCommand(conn *net.TCPConn) (ex error) {
	var init string
	conn.Write([]byte(START_COMM))
	buf := make([]byte, MAX_COMM_SIZE)

	_, err := conn.Read(buf)
	if (err != nil) {
		fmt.Fprintf(os.Stdout, "Error reading from buffer of TCP Connection\n%T\n%s\n", err, err)
		ex = err
		return
	}
	init = string(buf)
	if (init == RESP_OK) {
		jenc, _ := json.Marshal(comm)
		conn.Write([]byte(jenc))
        }else{
		ex = errors.New("Command was not accepted by remote server")
	}
	return ex
}
/*
ReceiveResp: receive the response from the remote server
*/
func (comm *TCPCommand) ReceiveResp(conn *net.TCPConn) (data *TCPExData, ex error){
	dt := make(chan []byte)
	errCh := make(chan error)
	var resp bytes.Buffer
	go func(){
		for {
			buf := make([]byte, MAX_BUFF_SIZE)
			_, err := conn.Read(buf)
			if err != nil {
				errCh<-err
			}
			dt<-buf
		}
	}()
	ex = <-errCh
	resp.Write(<-dt)
	ex = json.Unmarshal(resp.Bytes(), data)
	return
}
/*
SendData: send data on a previously opened connection; data is json-encoded
*/
func (data *TCPExData) SendData(conn *net.TCPConn) (ex error) {
	_, err := json.Marshal(data.Data)
	if err != nil {
		ex = err
	}
	//TODO
	return
}
