package tcpconn

const (
        PHASE_CONN = 0
        PHASE_AUTH = 1
        PHASE_SEND = 2
        PHASE_RETR = 3
        PHASE_SYNC = 4

        MAX_FILE_SIZE = 10485760
        MAX_BUFF_SIZE = 32768
        WRKR_COUNT = 100
        MAX_COMM_SIZE = 2048

	RESP_OK = "OK"
	START_COMM = "INIT"
)

type TCPCommand struct {
        OS, RHost, RPort string
        Phase int
}

type TCPExData struct {
	Data map[string]interface{}
}
