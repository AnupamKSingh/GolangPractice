package main
import	(
	"fmt"
_	"io"
_	"strconv"
	pb "google.golang.org/grpc/examples/oneone/oneone"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)
const	(
	address = "localhost:3030"
)
var (
	p2 *pb.Point2
	p1 *pb.Point1
)

func main()	{
	conn, err := grpc.Dial (address, grpc.WithInsecure())
	if err != nil {
		fmt.Println ("Error Happened in Dial")
	} else {
		fmt.Println ("Dial is passed")
	}
	clientcon := pb.NewOneoneClient(conn)
	
	p2, err = clientcon.GetFeature(context.Background(), &pb.Point1{A:12,B:23}); if err != nil {
		fmt.Println ("Error Happened in GetFeature")
	}	else	{
		fmt.Println ("Answersa are",p2.GetA(),p2.GetB())
	}
}
