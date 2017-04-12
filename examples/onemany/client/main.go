package main
import (
	"fmt"
	"io"
_	"strconv"
_	"net"
	
	pb "google.golang.org/grpc/examples/onemany/onemany"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
_	"google.golang.org/grpc/reflection"
)

const (
	address = "localhost:3030"
)

var rect *pb.Rect

func main ()	{
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println ("Error Happened")
	}	else	{
		clientconn := pb.NewOnemanyClient(conn)
		stream, err1 := clientconn.ReturnStream( context.Background(), & pb.Point{A:23,B:12}) 
		if err1 != nil {
			fmt.Println ("Error happened in ReturnStream")
		}	else	{
			for {
				rect, err = stream.Recv()
				if err == io.EOF	{
					fmt.Println ("EOF is Reached")
					break
				}	else	{
					if err != nil {
						fmt.Println ("Some unknown error happened")
					}	else	{
						fmt.Println (rect.GetA(), rect.GetB())
					}
				}
			}
		}
	}
}

		
