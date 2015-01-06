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
	fmt.Printf("starting connection\n")
	conn, err := comm.Connect()
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error initiating the connection\n%T\n%s", err, err)
	}
	fmt.Printf("posting command...\n")
	err = comm.PostCommand(conn)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error posting the command\n%T\n%s", err, err)
	}
	fmt.Printf("Waiting for response from server....")
	comm.ReceiveResp(conn)
	comm.Disconnect(conn)
	return
}

