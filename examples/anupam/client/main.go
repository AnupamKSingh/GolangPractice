package main
import (
	"fmt"
_	"strconv"
	"io"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pb "google.golang.org/grpc/examples/anupam/proto"
)

const (
	address = "localhost:50052"
)

func main()	{
	conn, err := grpc.Dial (address, grpc.WithInsecure())
    if  err != nil {
		fmt.Printf ("Error happened in grpc calls")
	}

	defer conn.Close()
	con := pb.NewAnupamClient (conn)
	if r, err := con.GetFeature(context.Background(), &pb.Point1{A:23,B:32});err != nil {
		fmt.Println("Error happened in GetFeature")
	} else {
		fmt.Println ("Result")
		fmt.Println(r.GetA().GetA(), r.GetB().GetA(),r.GetB().GetB(), r.GetA().GetB())
	}
	if stream, err := (pb.NewAnupamClient(conn)).SetFeature(context.Background(), &pb.Point2{A:21,B:12}); err != nil {
		fmt.Println("Error happened in Set Features")
	}	else	{
		for {
			RectStr , err := stream.Recv()
			if err == io.EOF {
				fmt.Println("End of io")
				break;
			}
			if err != nil	{
				fmt.Println ("Error")
				break;
			}
			fmt.Println (RectStr)
		}
	}
}
	

