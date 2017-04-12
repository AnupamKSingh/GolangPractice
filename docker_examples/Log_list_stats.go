package main

import (
    "context"
    "fmt"
	"io/ioutil"

    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
)

func main() {
    cli, err := client.NewEnvClient()
	fmt.Println ("Client = %s", cli)
    if err != nil {
		fmt.Println ("client.NewEnvClient")
        panic(err)
    }
    a := types.ContainerListOptions{}
    containers, err := cli.ContainerList(context.Background(), a)
    if err != nil {
		fmt.Println ("cli.ContainerList")
        panic(err)
    }

    for _, container := range containers {
        fmt.Printf("%s %s %d\n", container.ID[:10], container.Image, a.Limit)
		reader, _ := cli.ContainerLogs(context.Background(), container.ID, types.ContainerLogsOptions{
				ShowStdout: true,
				ShowStderr: true,
				Timestamps: true,
				Details:    true,
			})
		defer reader.Close()
		content, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Println ("Error Happened")
		}
		fmt.Println(string(content))
		stats, err := cli.ContainerStats (context.Background(), container.ID, false)
		if err != nil {
			fmt.Println ("Error Happened")
		}
		fmt.Println (stats.OSType)
		content1, err1 := ioutil.ReadAll(stats.Body)
        if err1 != nil {
            fmt.Println ("Error Happened")
        }
        fmt.Println(string(content1))

//		fmt.Println (stats.Size)
//		fmt.Println (stats.Mode)
	}


}

