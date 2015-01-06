package main

import(
	"flag"
	tcps "inge4pres/goutils/tcpconn"
)


func main(){

	proto := flag.String("proto", "tcp4", "Communication Protocol")
	port := flag.String("port", "4444", "Port for Communication")
	inet := flag.String("inet", "localhost","Interface to listen on")

	flag.Parse()

	tcps.TcpServerListener(*proto, *inet, *port)
}
