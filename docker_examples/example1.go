package main

import (
	"io"
	"os"
	"fmt"
	"time"

	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"golang.org/x/net/context"
)

func main() {
	image_name := "busybox:glibc"
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	fmt.Println ("Client Connection is done")
	_, err = cli.ImagePull(ctx, image_name, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println ("Image Pull is done")
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: image_name,
		Cmd:   []string{"top"},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	fmt.Println ("Container Creation is done")
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println ("Container start is done")
/*	if _, err = cli.ContainerWait(ctx, resp.ID); err != nil {
		panic(err)
	}
	fmt.Println ("Container Wait is done")
*/
	time.Sleep (10 * time.Second)
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)
}
