package ipc

import ( 
	"fmt"
	"encoding/json"
)

type IpcClient struct {
	conn chan string
}


func NewIpcClient( s *IpcServer ) *IpcClient {
	c := s.Connect()
	
	return &IpcClient{c}
}

func (client *IpcClient)Call( name, params string ) ( resp *Response, err error ) {
	req := &Request{ name, params }
	
	var b[] byte
	b, err = json.Marshal( req )
	if err != nil {
		fmt.Println( "Marshal request to json fail,", req )
		return
	}
	
	client.conn <- string(b)
	
	str := <- client.conn
	
	var resp1 Response
	err = json.Unmarshal( []byte(str), &resp1 )
	resp = &resp1
	
	return
}

func (client *IpcClient) Close() {
	client.conn <- "CLOSE"
}