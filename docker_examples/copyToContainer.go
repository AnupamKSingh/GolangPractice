package main

import (
    "context"
    "fmt"
	"os"
	"bufio"

    "github.com/docker/docker/client"
 	"github.com/docker/docker/api/types"
 	"github.com/docker/docker/api/types/container"
)

func main() {
    cli, err := client.NewEnvClient()
    fmt.Println ("Client = %s", cli)
    if err != nil {
        fmt.Println ("client.NewEnvClient")
        panic(err)
    }
	containerBody, err := cli.ContainerCreate (context.Background(), &container.Config{
				 Image: "816030111358",
				 Cmd:[]string{0:"/bin/sh",1:"/infiniteloop"},
				},
				 nil,nil, "anupam")
	if err != nil {
		fmt.Println (err)
	}
	fmt.Println (containerBody)
	err = cli.ContainerStart (context.Background(),"anupam", types.ContainerStartOptions {})
	if err != nil {
         fmt.Println (err)
    }
	os.Exit (0)
	Change, err := cli.ContainerDiff (context.Background(), "anupam")
	if err != nil {
         fmt.Println (err)
    }
    fmt.Println (Change)
	file, err := os.Open("dee.txt")
	if err != nil {
		 fmt.Println ("Open")
         fmt.Println (err)
    }
	b1 := make([]byte, 50)
    _, err = file.Read(b1)
	fmt.Println (b1)
	_, err = file.Seek(0, 0)
	r4 := bufio.NewReader(file)
    b4, err := r4.Peek(50)
	fmt.Printf("5 bytes: %s\n", string(b4))

	
	err = cli.CopyToContainer(context.Background(), "anupam", "/anu", bufio.NewReader(file), types.CopyToContainerOptions{
    	AllowOverwriteDirWithFile: false,
}) 
	if err != nil {
		 fmt.Println ("Copy")
         fmt.Println (err)
    }
	Change, err = cli.ContainerDiff (context.Background(), "anupam")
	if err != nil {
         fmt.Println (err)
    }
	
    fmt.Println (Change)
	  
}

