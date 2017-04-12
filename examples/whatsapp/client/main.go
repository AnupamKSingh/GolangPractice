package main

import (
	"fmt"
_	"io"
_	"io/ioutil"
_	"encoding/json"
_	"time"
_	"net"
_	"strconv"
	"bufio"
	"os"

	pb "google.golang.org/grpc/examples/whatsapp/whatsapp"
	"google.golang.org/grpc"
_	"google.golang.org/grpc/reflection"
	"golang.org/x/net/context"
)

const (
	address = "localhost:3023"
	yellow = "\x1b[32;1m"
	end = "\x1b[0m"
)

var Message *pb.Message

func main ()	{
	ch := make (chan int)
	conn, _ := grpc.Dial (address, grpc.WithInsecure())
	Client := pb.NewWhatsappClient (conn)
	stream, _ := Client.Chat (context.Background())
	go func ()	{
		for {
			message, err := stream.Recv()
			if err != nil || len(message.GetChat()) == 0{
				continue
			}
//			fmt.Println (len(message.GetChat()))
			fmt.Printf ("%s%s%s", yellow,message.Chat, end)
//			time.Sleep (2 * time.Second)
		}
	}()
	go func () {
		Message = new (pb.Message)
		for {
			Message.Chat = ""
			//fmt.Println ("Waiting for I/P")
			reader := bufio.NewReader (os.Stdin)
			Message.Chat, _ = reader.ReadString('\n')
//			fmt.Scanf ("%q", &Message.Chat)
			err := stream.Send (Message)
			if err != nil || len(Message.GetChat()) == 0{
				fmt.Println ("Error Starts")
				fmt.Println (Message.Chat)
				fmt.Println (err)
				fmt.Println ("Error Ends")
				continue
			}
//			time.Sleep (2 * time.Second)
		}
	}()
	<-ch
}
