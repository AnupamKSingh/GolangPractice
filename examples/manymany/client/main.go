package main
import	(
	"io"
	"fmt"
_	"strconv"
_	"net"
_	"encoding/json"

	pb "google.golang.org/grpc/examples/manymany/manymany"
	"google.golang.org/grpc"
_	"google.golang.org/grpc/reflection"
	"golang.org/x/net/context"
)

const	(
	address = "localhost:3030"
)

func main()	{
	conn, _ := grpc.Dial (address, grpc.WithInsecure())
	clientcon := pb.NewManymanyClient(conn)
	stream , _ := clientcon.ServerSendMany(context.Background(), &pb.Point{A:1,B:2})
	for {
		rect ,err := stream.Recv()
		if err == io.EOF	{
			fmt.Println ("EOF is reached")
			break
		}
		fmt.Println (rect.GetA().GetA(), rect.GetA().GetB(), rect.GetB().GetA(), rect.GetB().GetB())
	}
}
