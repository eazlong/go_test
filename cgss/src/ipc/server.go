package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

type Server interface {
	Name() string
	Handle( method, params string ) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer( server Server ) *IpcServer {
	return &IpcServer{ server } 
}

func (ipc *IpcServer)Connect() chan string {
	session := make( chan string, 0 )
	
	go func( c chan string ) {
		for {
			data := <- c
			
			if data == "CLOSE" {
				break
			}
			
			var req Request
			err := json.Unmarshal( []byte(data), &req )
			if err != nil {
				fmt.Println( "Invalid Request format", data );
				continue;
			}
			
			resp := ipc.Handle( req.Method, req.Params )
			b, err := json.Marshal( resp )
			
			c <- string(b)
			
			fmt.Println( "Response", resp )
			
		} 
	}( session )
	
	return session
}