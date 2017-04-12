package main
import (
_	"io"
	"fmt"
_	"io/ioutil"
_	"bufio"
	"net"
_	"syscall"
_	"time"
_	"strconv"
_	"encoding/json"

	pb "google.golang.org/grpc/examples/ping/ping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
_	"golang.org/x/net/context"
)

type PingServer struct {}

const (
	address = ":8081"
)

func (*PingServer) Pingpong (stream pb.Ping_PingpongServer) (error)	{
	fmt.Println ("I am inside recv")
	packet, _ := stream.Recv()
	fmt.Println ("I am outside recv")
	fmt.Println (packet)
	return nil
}

func newserver() (*PingServer)	{
	pingserver := new (PingServer)
	return pingserver
}

func main()	{
	lis, _ := net.Listen ("tcp", address)
	s := grpc.NewServer()
	pb.RegisterPingServer(s, newserver())
	reflection.Register(s)
	s.Serve(lis)
}
