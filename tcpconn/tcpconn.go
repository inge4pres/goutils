package tcpconn

import (
	"runtime"
	)
const (
        PHASE_CONN = 0 //Initial connection verification
        PHASE_AUTH = 1 //Authentication between client and server
        PHASE_SEND = 2 //Send data to server
        PHASE_RETR = 3 //Receive data from server
        PHASE_SYNC = 4 //Sync (optional)

        MAX_FILE_SIZE = 10485760  //10MB 
        MAX_BUFF_SIZE = 32768     //32KB
        WRKR_COUNT = 100          //Max number of workers for concurrency
        MAX_COMM_SIZE = 2048      //Max size of Commands JSON encoded

	RESP_OK = "OK"            //Default initial response, used for connectivity validation
	START_COMM = "INIT"       //Default initial request, used or connection validation
)

type TCPCommand struct {
        OS, LHost, LAddr, RAddr, RPort, Proto string
        Phase int
}

type TCPExData struct {
	Data map[string]interface{}
}

/*
init the TCPCOmmand type
*/
func init() {
	
}
