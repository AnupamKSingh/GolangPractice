package main
import	(
	"fmt"
_	"io"
_	"strconv"
	"net"
	pb "google.golang.org/grpc/examples/oneone/oneone"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	address = ":3030"
)


type OneoneServer struct {}

func newserver () *OneoneServer {
	s := new (OneoneServer)
	fmt.Println ("NewServer is being called")
	return s
}

func (s *OneoneServer) GetFeature(c context.Context, p1 *pb.Point1)	(p2 *pb.Point2, err error)	{
	fmt.Println ("Get Feature is being called")
	fmt.Println (p1.GetA(), p1.GetB())
	p2 = new (pb.Point2)
	p2.A = p1.A * 2
	p2.B = p1.B * 2
	fmt.Println (p2.GetA(), p2.GetB())
	return p2, nil
}
func main()	{
	lis, err := net.Listen ("tcp", address)
	if err != nil {
		fmt.Println ("Error happened in listen")
	}
	s := grpc.NewServer()
	pb.RegisterOneoneServer (s, newserver ())
	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		fmt.Println("Error Happened while serving")
	}
}
	
