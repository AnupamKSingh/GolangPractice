package main

import (
    "context"
    "fmt"

    "github.com/docker/docker/client"
 _	"github.com/docker/docker/api/types"
 _	"github.com/docker/docker/api/types/container"
)

func main() {
    cli, err := client.NewEnvClient()
    fmt.Println ("Client = %s", cli)
    if err != nil {
        fmt.Println ("client.NewEnvClient")
        panic(err)
    }
/*	containerBody, err := cli.ContainerCreate (context.Background(), &container.Config{Image: "busybox:glibc"}, nil,nil, "anupam3")
	if err != nil {
		fmt.Println (err)
	}
	fmt.Println (containerBody)
	err = cli.ContainerStart (context.Background(),"tender_hodgkin", types.ContainerStartOptions {})
	if err != nil {
         fmt.Println (err)
    }
*/	Change, err := cli.ContainerDiff (context.Background(), "jovial_jang")
	if err != nil {
         fmt.Println (err)
    }
    fmt.Println (Change)
	
	  
}

