package main
import	(
	"fmt"
_	"io"
	"net"
	"time"
_	"encoding/json"
_	"strconv"
_	"io/ioutil"

	pb "google.golang.org/grpc/examples/New1/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
_	"golang.org/x/net/context"
)

const	(
	address = ":3030"
)

type New1Server struct {}

//var message *pb.Message
var Message *pb.Message
var EnumMess *pb.EnumMessage

func (*New1Server) ServertoClient (stream pb.New1_ServertoClientServer) (error)	{
	ch := make (chan int)
	fmt.Println ("You got something from Server")
	go func ()	{
		fmt.Println ("Goroutine One started")
		for {
			message, err := stream.Recv ()
			if err != nil {
				continue
			}
			fmt.Println ("Recieve is Done")
			fmt.Println (message.A)
			fmt.Println (message.B)
			fmt.Println (message.C)
			fmt.Println (message.D)
			for bol := range message.E {
				fmt.Println (bol)
			}
			fmt.Println (message.F)
			fmt.Println (message.G)
			for str := range message.H {
				fmt.Println (str)
			}
			time.Sleep (1 * time.Millisecond)
		}
	}()
	go func ()	{
		for {
			fmt.Println ("Goroutine Two started")
			EnumMess = new (pb.EnumMessage)
	        EnumMess.A = 1
		    Message = new (pb.Message)
	        Message.A = 2000000;
		    Message.B = 200;
	        Message.C = 200.2000;
		    Message.D = "anupam kumar singh"
	        Message.E = []bool{false}
		    Message.F = EnumMess
	        Message.G = []byte("deepoo")
		    Message.H = []string{"anupam", "kumar", "singh"}
			stream.Send (Message)
			fmt.Println ("Send is Done")
	        time.Sleep (1 * time.Millisecond)
		}
	}()
	<-ch
	return nil
}
func  newserver() (*New1Server)	{
	newserve := new (New1Server)
	return newserve
}

func main ()	{
	lis, _ := net.Listen ("tcp", address)
	s := grpc.NewServer()
	pb.RegisterNew1Server (s, newserver())
	reflection.Register(s)
	s.Serve (lis)
}
