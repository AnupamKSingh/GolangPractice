package main;
import (
	"fmt"
	"io"
	"io/ioutil"
_	"net"
_	"strconv"
	"encoding/json"

	pb "google.golang.org/grpc/examples/allinone/allinone"
	"google.golang.org/grpc"
_	"google.golang.org/grpc/reflection"
	"golang.org/x/net/context"
)



const (
	address = "localhost:3030"
	filePath = "./client.json"
)

var savedfeature []*pb.Rect

func loadfeature (filepath string)	(error)	{
	fmt.Println ("Loading Json File")
	file, err := ioutil.ReadFile (filepath)
	if err != nil	{
		fmt.Println("Error happened in Reading the Client Json")
	}	else	{
		err = json.Unmarshal (file, &savedfeature)
		if err != nil	{
			fmt.Println ("Error happened in json.Unmarshal")
		}
	}
	return nil
}

func main()	{
	conn, _ := grpc.Dial (address, grpc.WithInsecure())
	Clientconn := pb.NewAllinoneClient (conn)

	loadfeature(filePath)	

	point, err := Clientconn.SendonetoServer (context.Background(), &pb.Point{A:2,B:4})
	if err != nil	{
		fmt.Println ("Error while SendonetoServer")
	}	else	{
		fmt.Println (point.GetA(), point.GetB())
	}
	
	stream, err1 := Clientconn.SendmanytoServer(context.Background())
	if err1 != nil	{
		fmt.Println ("Error Happenned in SendmanytoServer")
	}	else	{
		for _, one := range savedfeature	{
			err = stream.Send (one)
			if err != nil	{
				fmt.Println ("Error happened in Send of SendmanytoServer")
			}
		}
		fmt.Println ("Send completed in Client")
		point, err = stream.CloseAndRecv()
		if err != nil	{
			fmt.Println ("Error while CloseAndRecv() in SendmanytoServer")
		}	else	{
			fmt.Println ("What we got from the server")
			fmt.Println (point.GetA(), point.GetB())
		}
	}

	var rect *pb.Rect
	stream1 , err2 := Clientconn.SendmanytoClient (context.Background(), &pb.Point {A:1, B:2})
	if err2 != nil	{
		fmt.Println ("Error happenned in SendmanytoClient")
	}	else	{
		for {
			rect, err = stream1.Recv()
			if err == io.EOF	{
				fmt.Println ("EOF is reached")
				break
			}
			if err != nil	{
				fmt.Println ("Error happened in Recv of SendmanytoClient")
				break
			}	else	{
				fmt.Println (rect.GetA().GetA(), rect.GetA().GetB())
				fmt.Println (rect.GetB().GetA(), rect.GetB().GetB())
			}
		}
		
	}

}	
			
	
