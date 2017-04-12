package main;
import	(
	"fmt"
	"io"
_	"io/ioutil"
_	"encoding/json"
_	"strconv"
	"net"

	pb "google.golang.org/grpc/examples/manyone/manyone"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
_	"golang.org/x/net/context"
)

const	(
	address = ":3030"
)

type ManyoneServer struct{}

func (*ManyoneServer) MultiSendClient(stream pb.Manyone_MultiSendClientServer)	(error) {
	rect := new (pb.Rect)
	rect.A = new(pb.Point)
	rect.B = new(pb.Point)
	rect.B.A = 1
	rect.B.B = 1
	for {
		point ,err := stream.Recv()
		fmt.Println (err)
		if err == io.EOF	{
			fmt.Println ("Error Happened in Recv")
			break
		}
		fmt.Println(point.GetA(), point.GetB())
		rect.A.A += point.GetA()
		rect.A.B += point.GetB()
		rect.B.A *= point.GetA()
		rect.B.B *= point.GetB()
	}
	stream.SendAndClose(rect)
	return nil
}

func newserver ()	(*ManyoneServer)	{
	fmt.Println ("newserver is called")
	s := new (ManyoneServer)
	return s
}

func main()	{
	lis ,err := net.Listen("tcp", address)
	if err != nil	{
		fmt.Println ("listen api went wrong")
	}
	s := grpc.NewServer()
	fmt.Println ("NewServer got")
	pb.RegisterManyoneServer (s, newserver())
	fmt.Println ("NewServer got1")
	reflection.Register(s)
	fmt.Println ("NewServer got2")
	s.Serve(lis)
}

