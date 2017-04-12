package main
import (
	"fmt"
	"net"
	"encoding/json"
_	"io"
	"io/ioutil"
_	"strconv"
	
	pb "google.golang.org/grpc/examples/onemany/onemany"
_	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const	(
	address = ":3030"
	filepath = "./onemany.json"
)
var savedfeatures []*pb.Rect
type OnemanyServer struct{}

func (s *OnemanyServer) loadfeatures(filepath string)	{
	file ,err := ioutil.ReadFile(filepath)
	if err != nil	{
		fmt.Println ("Error happened while doing read")
	}
	err = json.Unmarshal(file, &savedfeatures)
	if err != nil	{
		fmt.Println ("Error happened while doing unmarshall")
	}
}


func (s *OnemanyServer)	ReturnStream (p *pb.Point, stream pb.Onemany_ReturnStreamServer) (error)	{	
	for _, feature := range savedfeatures	{
		err := stream.Send(feature)
		if err != nil	{
			fmt.Println ("Error happened while sending the data")
			return err
		}
	}
	return nil
}
	

func  newserver() (*OnemanyServer)	{
	s := new(OnemanyServer)
	s.loadfeatures(filepath)
	return s
}

func main ()	{
	lis, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println ("Error Happened")
	}	
	s := grpc.NewServer()
	pb.RegisterOnemanyServer(s, newserver())
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Println ("Error happened while servinhg from Server")
	}
}
	
