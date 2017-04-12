package main
import (
_	"io"
_	"io/ioutil"
_	"os"
_	"bufio"
	"fmt"
	"time"
_	"bytes"
	"reflect"
_	"encoding/json"

	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"golang.org/x/net/context"
)

var netID string
func main()	{
	image_name := "outyet"
	cont_name := "anupam"
	new_cont_name := "deepoo1"
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil	{
		panic (err)
	} else {
		fmt.Println ("Image pull is done")
	}

	version := cli.ClientVersion ()
	fmt.Println ("Client Version used :", version)

//	Working on local images
/*	_, err = cli.ImagePull (ctx, image_name, types.ImagePullOptions{})
	if err != nil	{
		panic (err)
	} else {
		fmt.Println ("Image Pull completed")
	}
*/
	contBody, err := cli.ContainerCreate (ctx, &container.Config{
				Image : image_name,
/*				Cmd : []string {"top"},*/
				ExposedPorts: nat.PortSet{
					"8080/tcp": {},
					},
				},
				&container.HostConfig{
					AutoRemove :true,
					Privileged : true,
					//PortBindings:map[string]string{
					PortBindings:nat.PortMap{
								"8080/tcp":[]nat.PortBinding {
									nat.PortBinding{HostPort:"6162"},},
							},
							},
				nil , cont_name)
	if err!= nil	{
		panic(err)
	} else {
		fmt.Println ("Container Created")
	}

	err = cli.ContainerStart (ctx, contBody.ID, types.ContainerStartOptions{})
	if err != nil	{
		panic (err)
	} else {
		fmt.Println ("Container Started")
	}

	pathStat, errp := cli.ContainerStatPath (ctx, contBody.ID, "/")
	if errp != nil {
		panic (err)
	} else {
		fmt.Println ("ContainerStatPath Starts")
		fmt.Println (pathStat.Name)
		fmt.Println (pathStat.Size)
		fmt.Println (pathStat.Mode)
		fmt.Println (pathStat.Mtime)
		fmt.Println (pathStat.LinkTarget)
		fmt.Println ("ContainerStatPath Stops")
	}
/*
	ctop, errt := cli.ContainerTop (ctx, contBody.ID, []string {})
	if errt != nil {
		panic (errt)
	} else {
		fmt.Println ("ContainerTop Starts")
		fmt.Println (ctop.Processes)
		fmt.Println (ctop.Titles)
		fmt.Println ("ContainerTop Stops")
	}

	ioread, cpcon, errcp := cli.CopyFromContainer (ctx, contBody.ID, "/anu")
	if errcp != nil {
		panic (errcp)
	} else {
		fmt.Println ("CopyFromContainer Starts")
		buf := new (bytes.Buffer)
		buf.ReadFrom (ioread)
		fmt.Println (buf.String())
		fmt.Println (ioutil.ReadAll(ioread))
		fmt.Println (cpcon.Name)
		fmt.Println (cpcon.Size)
		fmt.Println (cpcon.Mode)
		fmt.Println (cpcon.Mtime)
		fmt.Println (cpcon.LinkTarget)
		fmt.Println ("CopyFromContainer Stops")
	}
*/
	diskuse, erru := cli.DiskUsage (ctx)
	if erru != nil {
		panic (err)
	} else {
		fmt.Println ("DiskUsage Starts")
		fmt.Println ("DiskUsage Image Starts")
		fmt.Println (diskuse.LayersSize)
		fmt.Println (diskuse.Images[0].Containers)
		fmt.Println (diskuse.Images[0].Created)
		fmt.Println (diskuse.Images[0].ID)
		fmt.Println (diskuse.Images[0].RepoTags)
		fmt.Println (diskuse.Images[0].SharedSize)
		fmt.Println (diskuse.Images[0].Size)
		fmt.Println (diskuse.Images[0].VirtualSize)
		fmt.Println ("DiskUsage Image Stops")
		fmt.Println (diskuse.Volumes)
		fmt.Println ("DiskUsage Container Starts")
		fmt.Println (diskuse.Containers[0].ID)
		fmt.Println (diskuse.Containers[0].Names)
		fmt.Println (diskuse.Containers[0].Image)
		fmt.Println (diskuse.Containers[0].ImageID)
		fmt.Println (diskuse.Containers[0].Command)
		fmt.Println (diskuse.Containers[0].Created)
		fmt.Println (diskuse.Containers[0].Ports)
		fmt.Println (diskuse.Containers[0].SizeRw)
		fmt.Println (diskuse.Containers[0].SizeRootFs)
		fmt.Println (diskuse.Containers[0].State)
		fmt.Println (diskuse.Containers[0].Status)
		fmt.Println ("DiskUsage Container Stops")
	}

	imagehis, errih := cli.ImageHistory (ctx, "anupam:latest")
	if errih!= nil {
		panic (errih)
	} else {
		fmt.Println ("ImageHistory Starts")
		fmt.Println (imagehis[0].Comment)
		fmt.Println (imagehis[0].Created)
		fmt.Println (imagehis[0].CreatedBy)
		fmt.Println (imagehis[0].ID)
		fmt.Println (imagehis[0].Size)
		fmt.Println (imagehis[0].Tags)
		fmt.Println ("ImageHistory Stops")
	}
//	Some Unexpected EOF while doing tar comes for Copying to Contoiner
/*
	time.Sleep (20 * time.Second)
	file, errc := os.Open("./anupam")
	errc = cli.CopyToContainer(ctx, contBody.ID, "./", bufio.NewReader(file), types.CopyToContainerOptions{
		AllowOverwriteDirWithFile: false,
	})
	if errc!= nil	{
		panic (errc)
	} else {
		fmt.Println ("CopyToContainer is done")
	}

*/
//	time.Sleep (10 * time.Second)

//	out, err1 := cli.ContainerLogs (ctx, contBody.ID, types.ContainerLogsOptions {
	_, err1 := cli.ContainerLogs (ctx, contBody.ID, types.ContainerLogsOptions {
				ShowStdout: true,
				ShowStderr: true,
				Timestamps: true,
                Details:    true,})
	if err1!= nil	{
		panic (err1)
	} else {
		fmt.Println ("Got the container logs")
	}

	cont_list, err2 := cli.ContainerList (ctx, types.ContainerListOptions {
				Size: true,
				All : true,})
	if err2 != nil	{
		panic (err2)
	} else {
		fmt.Println ("Container List is Obtained")
		for _, cont := range cont_list {
			fmt.Println (cont.ID)
			fmt.Println (cont.Names)
			fmt.Println (cont.Image)
			fmt.Println (cont.ImageID)
			fmt.Println (cont.Created)
			fmt.Println (cont.Command)
			fmt.Println (cont.SizeRw)
			fmt.Println (cont.SizeRootFs)
		}
	}

	err = cli.ContainerRename (ctx, contBody.ID, new_cont_name)
	if err != nil	{
		panic (err)
	} else {
	fmt.Println ("Container Name is changed")
	}


	json1, errj := cli.ContainerInspect (ctx, contBody.ID)
	if err != nil {
		panic (errj)
	} else {
		fmt.Println ("Container Inspect is passed")
		fmt.Println (json1.Config)
		fmt.Println (json1.NetworkSettings)
		fmt.Println (json1.ContainerJSONBase)
	}

/*	errkill := cli.ContainerKill (ctx, contBody.ID, "SIGKILL")
	if errkill != nil	{
		panic (errkill)
	} else {
		fmt.Println ("Container Kill is done")
	}
*/
/*
    ipamcreate := new(network.IPAM)
	var ipamcreateConfig =  [...]network.IPAMConfig {0:network.IPAMConfig{Gateway:"23.23.23.1", Subnet:"23.23.23.0/24"}}
	ipamcreate.Config = ipamcreateConfig[:]
	netResp, errnet := cli.NetworkCreate (ctx, "newNet1", types.NetworkCreate {
												CheckDuplicate : true,
												IPAM :  ipamcreate,
												EnableIPv6     : false,
												Attachable     : true,
												Internal       : true,})
	if errnet!= nil {
		panic (errnet)
	} else {
		fmt.Println ("Network Create is Done")
		fmt.Println (netResp.ID)
		fmt.Println (netResp.Warning)
	}
*/
	nets , errnets := cli.NetworkList (ctx, types.NetworkListOptions {})
	if errnets != nil {
		panic (errnets)
	} else {
		fmt.Println ("Network List is obtained")
		for _, net := range nets {
				fmt.Println ("ID from Anupam: ", net.ID)
				fmt.Println (net.Name)
				for _, ipam := range net.IPAM.Config {
					fmt.Println (ipam.Gateway)
					fmt.Println (ipam.Subnet)
					fmt.Println (ipam.IPRange)
				}
				fmt.Println (net.Internal)
				fmt.Println (net.Attachable)
				fmt.Println (net.Driver)
		}
	}

	var config network.EndpointSettings
	config.IPAMConfig = new (network.EndpointIPAMConfig)
	config.IPAMConfig.IPv4Address = "23.23.23.2"
	config.NetworkID = "newNet1"
	config.IPAddress = "172.19.0.17"
	config.Gateway = "172.19.0.1"
	config.IPPrefixLen = 16
	errnetcon := cli.NetworkConnect (ctx, "newNet1", contBody.ID, &config)
	if errnetcon != nil {
		panic (errnetcon)
	} else {
		fmt.Println ("NetConnection Done")
	}

	resource, errres := cli.NetworkInspect (ctx, "newNet1")
	if errres != nil {
		panic (errres)
	}
	fmt.Println (resource.Name)
	fmt.Println (resource.ID)
//	fmt.Println (resource.Created)
	fmt.Println (resource.Scope)
	fmt.Println (resource.Driver)
	fmt.Println (resource.EnableIPv6)
	fmt.Println (resource.Internal)
	fmt.Println (resource.Attachable)
	fmt.Println (resource.IPAM.Driver)
	for _, config := range resource.IPAM.Config {
		fmt.Println (config.Subnet)
		fmt.Println (config.IPRange)
		fmt.Println (config.Gateway)
		fmt.Println (config.AuxAddress)
	}
	for _, cont := range resource.Containers {
		fmt.Println (cont.Name)
//		fmt.Println (cont.EndpoindID)
		fmt.Println (cont.MacAddress)
		fmt.Println (cont.IPv4Address)
		fmt.Println (cont.IPv6Address)
	}
/* in cli.Events we are hanging while reading from the channel */
	go func(){
		for {
			chanMess, _:= cli.Events (ctx, types.EventsOptions {})
			Mess := <-chanMess
			fmt.Println (reflect.TypeOf(Mess))
			fmt.Println ()
			fmt.Println ("Mess.Status")
			fmt.Println (Mess.Status)
			fmt.Println ()
			fmt.Println ("Mess.ID")
			fmt.Println (Mess.ID)
			fmt.Println ()
			fmt.Println ("Mess.From")
			fmt.Println (Mess.From)
			fmt.Println ()
			fmt.Println ("Mess.Type")
			fmt.Println (Mess.Type)
			fmt.Println ()
			fmt.Println ("Mess.Actor.ID")
			fmt.Println (Mess.Actor.ID)
			fmt.Println ()
			fmt.Println ("Mess.Actor.Attributes")
			fmt.Println (Mess.Actor.Attributes)
			fmt.Println ()
			fmt.Println ("Mess.Action")
			fmt.Println (Mess.Action)
			fmt.Println ()
			fmt.Println ("Mess.Time")
			fmt.Println (time.Unix(0,Mess.Time*int64(time.Millisecond)))
			fmt.Println ()
			fmt.Println ("Mess.TimeNano")
			fmt.Println (time.Unix(0,Mess.TimeNano*int64 (time.Millisecond)))
			fmt.Println ()

		}
	}()

	/* Error cam ewhen we tried to enable RemoveVolumes and RemoveLinks 
	panic: Error response from daemon: Conflict, cannot remove the default name of the container */
/*	errem := cli.ContainerRemove (ctx, contBody.ID, types.ContainerRemoveOptions {
								//	RemoveVolumes: true,
								//	RemoveLinks  : true,
									Force        : true,})
	if errem != nil {
			panic (errem)
	} else {
		fmt.Println ("Container is removed")
	}
*/
//	io.Copy (os.Stdout, out)

/*	err = cli.Close ()
	if err != nil	{
		panic (err)
	} else {
		fmt.Println ("Docker Client is closed successfully")
	}
*/
/*	hijackcon, errhi := cli.ContainerAttach (ctx, contBody.ID, types.ContainerAttachOptions{
													Stdout : true,
													Stdin  : true,
													Stderr : true,})
	if errhi != nil {
		panic (errhi)
	} else {
		fmt.Println ("Connection is Hijacked")
		go func ()	{
			for {
				tmp := make([]byte, 4096)
//				_, err := hijackcon.Reader.Read(tmp)
				_, err := hijackcon.Conn.Read(tmp)
				if err != nil {
					time.Sleep (1 * time.Second)
					fmt.Println ("Error in Attach:", err)
					continue
				} else {
					fmt.Println (tmp)
				}
			}
		}()
	}*/
	var chananu chan int
	<-chananu
}
