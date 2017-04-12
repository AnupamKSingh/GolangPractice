package main

import (
    "context"
    "fmt"

    "github.com/docker/docker/client"
    "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types"
)

func main() {
	cont_name := "Anupam"
    cli, err := client.NewEnvClient()
    fmt.Println ("Client = %s", cli)
    if err != nil {
        fmt.Println ("client.NewEnvClient")
        panic(err)
    }
	containerBody, err := cli.ContainerCreate (context.Background(), &container.Config{Image: "busybox:glibc"}, nil,nil, cont_name)
	if err != nil {
		fmt.Println (err)
	}
	fmt.Println (containerBody)
	err = cli.ContainerStart (context.Background(),cont_name, types.ContainerStartOptions {})
	if err != nil {
         fmt.Println (err)
    }  
/*	ec := types.ExecConfig{}
    ec.Detach = false
    ec.Tty = false
    ec.Cmd = make([]string, 3)
    ec.Cmd[0] = "/bin/bash"
    ec.Cmd[1] = "-c"
    ec.Cmd[2] = "ls"
	ec.AttachStdout = true	
	Response, err := cli.ContainerExecCreate (context.Background(), "tender_hodgkin", ec)
	if err !=nil {
		fmt.Println (err)
	}
	fmt.Println (Response)
	sc := types.ExecStartCheck{}
	sc.Tty = true
    sc.Detach = false
    err = cli.ContainerExecStart(context.Background(), Response.ID, sc)
    if err != nil {
		fmt.Println (err)
    }
	info, err := cli.ContainerExecInspect(context.Background(), Response.ID)
	if err !=nil {
		fmt.Println (err)
	} 
	fmt.Println (info)
*/
}

