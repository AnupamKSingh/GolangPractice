package main
import 	(
	"io"
	"io/ioutil"
	"fmt"
	"net"
	"encoding/json"
_	"strconv"

	pb "google.golang.org/grpc/examples/allinone/allinone"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

const	(
	address	= ":3030"
	filePath = "./server.json"
)

var savedfeature []*pb.Rect

type AllinoneServer struct {}

func loadfeature (filepath string)	(error)	{
	file, err := ioutil.ReadFile (filePath)
	if err != nil	{
		fmt.Println ("Error happened in ioutil.ReadFile")
		fmt.Println (err)
		return err
	}
	err = json.Unmarshal (file, &savedfeature)
	if err != nil	{
		fmt.Println ("Error happened in json.Unmarshall")
		fmt.Println (err)
		return err
	}
	return nil
}
func (s * AllinoneServer) SendonetoClient (c context.Context, point *pb.Point)	(*pb.Point, error)	{
	fmt.Println ("Recieved from Client in one-to-one communication")
	fmt.Println (point.GetA(), point.GetB())
	p := new (pb.Point)
	p.A = point.GetA() * 2
	p.B = point.GetB() * 2
	return p, nil
}

func (s * AllinoneServer) SendonetoServer (c context.Context, point *pb.Point)	(*pb.Point, error)	{
	fmt.Println ("Recieved from Client in one-to-one communication")
	fmt.Println (point.GetA(), point.GetB())
	p := new (pb.Point)
	p.A = point.GetA() * 2
	p.B = point.GetB() * 2
	return p, nil
}

func (s *AllinoneServer) SendmanytoClient (point *pb.Point, stream pb.Allinone_SendmanytoClientServer)	(error)	{
	fmt.Println ("Recieved from Client in one-to many conversation")
	fmt.Println (point.GetA(), point.GetB())
	for _, rect := range savedfeature	{
		err := stream.Send(rect)
		if err != nil	{
			fmt.Println ("Error happened in Send()")
			return err
		}
	}
	return nil
}

func (s * AllinoneServer) SendmanytoServer (stream pb.Allinone_SendmanytoServerServer)	(error)	{
	fmt.Println ("We are going to recv many from Client")
	point := new (pb.Point)
	for	{
		rect, err := stream.Recv()
		if err == io.EOF	{
			fmt.Println ("EOF is reached")
			break
		}
		if err != nil	{
			fmt.Println ("Unknown error while recieving from client")
			break
		}	else	{
			fmt.Println (rect.GetA().GetA(), rect.GetA().GetB())
			fmt.Println (rect.GetB().GetA(), rect.GetB().GetB())
			point.A += (rect.GetA().GetA() + rect.GetB().GetA())
			point.B += (rect.GetA().GetB() + rect.GetB().GetB())
		}
	}
	_ = stream.SendAndClose (point)
	return nil
}


func newserver () (*AllinoneServer)	{
	s := new (AllinoneServer)
	loadfeature(filePath)
	return s
}

func main()	{
	lis, err := net.Listen ("tcp", address)
	if err != nil	{
		fmt.Println ("Error happened in net.Listen")
	}
	s := grpc.NewServer()
	pb.RegisterAllinoneServer (s, newserver())
	reflection.Register(s)
	s.Serve(lis)
}

