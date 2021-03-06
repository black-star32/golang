package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRPC(host string, service interface{}) error{
	rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err != nil{
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil{
			log.Printf("accept error: %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
	return nil
}

func NewClient(host string)(*rpc.Client, error){
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
