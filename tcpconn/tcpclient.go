package tcpconn

import (
	"bytes"
	json "encoding/json"
	"net"
)
/*
Connect: starts a IPv4 connection with a remote server with protocol Command.Proto, leaves it open for future use
*/
func (comm *TCPCommand) Connect() (conn *net.TCPConn, err error) {
	addr, err := net.ResolveTCPAddr(comm.Proto, comm.RAddr+":"+comm.RPort)
	conn, err = net.DialTCP(comm.Proto, nil, addr)
	if err != nil {
		e.HandlErr("FATAL ", "Could not reach the remote server at "+comm.RAddr+":"+comm.RPort+"\n", err)
	}
	return conn, err
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
func (comm *TCPCommand) PostCommand(conn *net.TCPConn) error {
	jc, err := json.Marshal(comm)
	if err == nil {
		conn.Write(jc)
	} else {
		e.HandlErr("FATAL", "Could not post command: json.Marshal() failed", err)
	}
	return err
}

/*
ReceiveResp: receive the response from the remote server
*/
func (comm *TCPCommand) ReceiveResp(conn *net.TCPConn) (data *TCPExData, err error) {
	dt := make(chan []byte)
	errCh := make(chan error)
	var resp bytes.Buffer
	go func(){
		buf := make([]byte, MAX_BUFF_SIZE)
		b, err := conn.Read(buf)
		if err != nil {
			errCh <- err
		}
		dt <- buf[0:b]
	}()
	err = <-errCh
	if err != nil{
		e.HandlErr("WARN", "Error reading the response from remote server", err)
	}
	resp.Write(<-dt)
	data = &TCPExData{}
	err = json.Unmarshal(resp.Bytes(), data)
	conn.Close()
	return
}

/*
SendData: send data on a previously opened connection; data is json-encoded
*/
func (data *TCPExData) SendData(conn *net.TCPConn) error {
	denc, err := json.Marshal(data.Data)
	if err != nil {
		e.HandlErr("WARN", "The data format is incorrect!", err)
		return err
	}
	conn.Write(denc)
	return err
}
