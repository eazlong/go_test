package ipc

import "testing"
import "fmt"

type EchoServer struct {
}

func ( s *EchoServer ) Handle( method, params string ) *Response{
	var resp Response;
	resp.Code = "200";
	resp.Body = params;
	
	fmt.Println( "receive methed", method, "params", params, "return", resp.Code )
	
	return &resp;
}

func ( s *EchoServer ) Name() string {
	return "EchoServer"
}

func testIpc( t *testing.T ) {
	ipcServer := NewIpcServer( &EchoServer{} )
	ipcServer.Connect()
	
	c1 := NewIpcClient( ipcServer )
	c2 := NewIpcClient( ipcServer )
	
	resp1, err1 := c1.Call( "call", "From c1" )
	resp2, err2 := c2.Call( "call", "From c2" )
	
	t.Error( resp1.Code, ":", resp1.Body, resp2.Code, ":", resp2.Body )
	
	if err1 != nil || err2 != nil {
		t.Error( err1, err2 )
	}
	c1.Close()
	c2.Close()
}