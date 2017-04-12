package main;

import (
	"fmt"
_	"io"
_	"encoding/json"
	"time"
_	"net"

	pb "google.golang.org/grpc/examples/New1/proto"
	"google.golang.org/grpc"
_	"google.golang.org/grpc/reflection"
	"golang.org/x/net/context"
)

const (
	address = "localhost:3030"
)

var Message *pb.Message
//var Rmessage *pb.Message
var EnumMess *pb.EnumMessage

func main ()	{
	ch := make (chan int)
	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	ClientConn := pb.NewNew1Client(conn)
	stream, _ := ClientConn.ServertoClient (context.Background())
	go func () {
		for {
			EnumMess = new (pb.EnumMessage)
			EnumMess.A = 0
			Message = new (pb.Message)
			Message.A = 1000000;
			Message.B = 100;
			Message.C = 100.1000;
			Message.D = "Anupam Kumar Singh"
			Message.E = []bool{true}
			Message.F = EnumMess
			Message.G = []byte("Deepoo")
			Message.H = []string{"Anupam", "Kumar", "Singh"}
			stream.Send (Message)
			time.Sleep (1 * time.Millisecond)
		}
	}()

	go func ()	{
		for {
			Rmessage, err := stream.Recv ()
			if err != nil {
				continue
			}
			fmt.Println (Rmessage.GetA())
			fmt.Println (Rmessage.GetB())
			fmt.Println (Rmessage.GetC())
			fmt.Println (Rmessage.GetD())
			fmt.Println (Rmessage.GetE())
			fmt.Println (Rmessage.GetF())
			fmt.Println (Rmessage.GetG())
			for str := range Rmessage.GetH() {
					fmt.Println (str)
			}
			time.Sleep (1 * time.Millisecond)
		}
	}()
	<-ch
}
