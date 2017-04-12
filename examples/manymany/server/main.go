package main
import (
_	"io"
	"io/ioutil"
	"fmt"
	"net"
_	"strconv"
	"encoding/json"
	
	pb "google.golang.org/grpc/examples/manymany/manymany"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
_	"golang.org/x/net/context"
)

const	(
	address = ":3030"
	filepath = "./server.json"
)

type ManymanyServer struct{}

var savedfeature []*pb.Rect

func loadfeature (filePath string)	{
	file, err := ioutil.ReadFile(filePath)
	if err != nil	{
		fmt.Println("Error while ioutil.ReadFile")
	}
	err = json.Unmarshal(file, &savedfeature)
	if err != nil {
		fmt.Println("Error happened in json.Unmarshal")
	}
}

func (*ManymanyServer)	ServerSendMany(p *pb.Point, stream pb.Manymany_ServerSendManyServer)	(error){
	fmt.Println ("What we got from Client is")
	fmt.Println (p.GetA(), p.GetB())
	rect := new(pb.Rect)
	rect.A = new (pb.Point)
	rect.B = new (pb.Point)
	for	i := 0;i <5; i ++	{
		rect.A.A = p.GetA() * int32(i)
		rect.A.B = p.GetA() * int32(i + 1)
		rect.B.A = p.GetA() * int32(i + 2)
		rect.B.B = p.GetA() * int32(i + 3)
		err := stream.Send (rect)
		if err != nil 	{
			fmt.Println ("Error Happened in Send")
			return err
		}
	}
	return nil
}

func newserver()	(s *ManymanyServer)	{
	s = new(ManymanyServer)
	loadfeature(filepath)
	return s
}

func main ()	{
	lis, err := net.Listen("tcp", address)
	if err != nil	{
		fmt.Println ("Error Happened in Listen")
	}
	s := grpc.NewServer()
	pb.RegisterManymanyServer(s, newserver())
	reflection.Register(s)
	s.Serve(lis)
}	

	
