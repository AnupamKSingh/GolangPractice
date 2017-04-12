package main
import (
	"fmt"
_	"strconv"
	"net"
	"io/ioutil"
	"encoding/json"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"golang.org/x/net/context"
	pb "google.golang.org/grpc/examples/anupam/proto"
)

type anupamserver struct{}

var savedfeature []*pb.RectStr

func (s *anupamserver) loadfeatures (filepath string) {
	file, err := ioutil.ReadFile (filepath)
	if err != nil	{
		fmt.Println("File not opened")
	}
	err = json.Unmarshal(file, &savedfeature)
	if err != nil {
		fmt.Println("Error happened in json Unmarshall")
	}
}

	
func newServer() (*anupamserver)	{
	s := new (anupamserver)
	s.loadfeatures (FilePath)
	return s
}

const (
	port = ":50052"
	FilePath = "./anupam.json"
)

func (s *anupamserver) GetFeature(c context.Context,p *pb.Point1) (*pb.Rect, error) {
	fmt.Println("We are in GetFeature")
	r := &pb.Rect{A:p,B:p}
	return r, nil
}

func (s *anupamserver) SetFeature(p *pb.Point2, stream pb.Anupam_SetFeatureServer) (error) {
	for _, feature := range savedfeature {
		if err := stream.Send(feature);err != nil {
			fmt.Println("Error in Send")
			return err
		}
	}
	return nil 
}

func main()	{
	lis, err := net.Listen("tcp", port)
	if  err != nil {
		fmt.Println ("Error happened in Listen")
	}
	s := grpc.NewServer()
	//pb.RegisterAnupamServer (s, &anupamserver{})
	pb.RegisterAnupamServer (s, newServer())
	reflection.Register(s)
	if err := s.Serve(lis);err != nil {
		fmt.Println ("Error in serving at the port")
	}
}
	
	
	
