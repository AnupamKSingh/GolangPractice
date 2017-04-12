package main

import (
	"fmt"
_	"io"
_	"io/ioutil"
	"net"
_	"time"
_	"encoding/json"
_	"strconv"
	"bufio"
	"os"

	pb "google.golang.org/grpc/examples/whatsapp/whatsapp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
_	"golang.org/x/net/context"
)

const (
	address = ":3023"
	yellow = "\x1b[32;1m"
	end = "\x1b[0m"
)

//Color examples
/*
const CLR_0 = "\x1b[30;1m"

const CLR_R = "\x1b[31;1m"

const CLR_G = "\x1b[32;1m"

const CLR_Y = "\x1b[33;1m"

const CLR_B = "\x1b[34;1m"

const CLR_M = "\x1b[35;1m"

const CLR_C = "\x1b[36;1m"

const CLR_W = "\x1b[37;1m"

const CLR_N = "\x1b[0m"

*/

type WhatsappServer struct {}

var Message *pb.Message

func newserver () (*WhatsappServer)	{
	whatsappserver := new (WhatsappServer)
	return whatsappserver
}

func (*WhatsappServer) Chat (stream pb.Whatsapp_ChatServer) (error)	{
	fmt.Println ("We are inside Server. Lets run forever")
	ch := make (chan int)
	go func ()	{
		for {
			message, err := stream.Recv ()
			if err != nil || len(message.GetChat()) == 0 {
				continue
			}
			fmt.Printf ("%s%s%s", yellow, message.GetChat(), end)
		}
	}()
	go func ()	{
		Message = new (pb.Message)
		for {
			Message.Chat = ""
//			fmt.Println ("Waiting for Input")
			reader := bufio.NewReader (os.Stdin)
			Message.Chat, _ = reader.ReadString ('\n')
//			fmt.Scanf ("%q",&Message.Chat)
			err := stream.Send (Message)
			if err != nil {
				continue
			}
//			time.Sleep (2 * time.Second)
		}
	}()
	<-ch
	return nil
}


func main ()	{
	lis, _ := net.Listen ("tcp", address)
	s := grpc.NewServer()
	pb.RegisterWhatsappServer (s, newserver())
	reflection.Register (s)
	s.Serve (lis)
}
