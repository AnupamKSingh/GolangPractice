package main;
import (
	"fmt"
	"io"
	"io/ioutil"
_	"net"
_	"strconv"
	"encoding/json"

	pb "google.golang.org/grpc/examples/manyone/manyone"
_	"google.golang.org/grpc/reflection"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var savedfeature []*pb.Point

const (
	address = "localhost:3030"
	filepath = "./manyone.json"
)

func loadFeature(filepath string)	{
	fmt.Println ("Load Feature is called")
	file, err := ioutil.ReadFile(filepath)
	if err != nil {	
		fmt.Println("Error in io/ioutil")
	}
	err = json.Unmarshal(file, &savedfeature)
	if err != nil {
		fmt.Println ("Error in json.Unmarshal")
	}
	fmt.Println (savedfeature)
}


func main()	{
	conn, err := grpc.Dial (address, grpc.WithInsecure())
	clientconn := pb.NewManyoneClient(conn)
	loadFeature(filepath)
	stream, _ := clientconn.MultiSendClient(context.Background())
	fmt.Println ("Got the stream now")
	for _, one := range savedfeature	{
		fmt.Println (one)
		err = stream.Send(one)	
		fmt.Println ("Send happened")
		if err == io.EOF	{
			fmt.Println ("EOF is Reached")
			break
		}
		if err != nil	{
			fmt.Println ("Error Happedned")
			break
		}
	}
	rect, _ := stream.CloseAndRecv()
	fmt.Println(rect.GetA().GetA(), rect.GetA().GetB(), rect.GetB().GetA(), rect.GetB().GetB())	
}
	
	
