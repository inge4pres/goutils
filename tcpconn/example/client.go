package main

import (
	tconn "inge4pres/goutils/tcpconn"
	"flag"
	"fmt"
	"os"
	"runtime"
)

func main(){

	host := flag.String("host", "localhost", "The remote server address")
	port := flag.String("port", "4444", "The remote server port")
	//path := flag.String("path", "", "The directory to check")
	proto := flag.String("proto", "tcp4", "The connection protocol")

	flag.Parse()

	comm := tconn.TCPCommand{Proto:*proto,OS:runtime.GOOS,RAddr:*host,RPort:*port,Phase:tconn.PHASE_CONN}
	conn, err := comm.Connect()
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error initiating the connection\n%T\n%s", err, err)
	}
	err = comm.PostCommand(conn)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error posting the command\n%T\n%s", err, err)
	}
	fmt.Println("Waiting for response from server....")
	data, err := comm.ReceiveResp(conn)
	if  err != nil {
		fmt.Println("Error receiving response: ")
		fmt.Println(err)
	}
	fmt.Println("DATA: ")
	fmt.Println(data)
	comm.Disconnect(conn)
	return
}

