package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:   "compiler",
		Cmd:     []string{"bash"},
		Tty:     true,
		Volumes: map[string]struct{}{"./vol/": {}},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}
	fmt.Print(resp.Warnings)
	fmt.Print(resp.ID)
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	fmt.Print("Started Container")

	/*_, runit := exec.Command("/bin/bash", "-c", "docker start "+resp.ID[:12]).Output()
	if runit != nil {
		panic(runit)
	}*/

	con := types.ExecConfig{
		Cmd: []string{"gcc", "/vol/main.c -o main"},
	}

	execID, err := cli.ContainerExecCreate(ctx, resp.ID, con)
	if err != nil {
		panic(err)
	}
	fmt.Println("ExecProcess Created")
	cli.ContainerExecStart(ctx, execID.ID, types.ExecStartCheck{})

	con2 := types.ExecConfig{
		Cmd: []string{"./vol/main > data.txt"},
	}

	execID2, err := cli.ContainerExecCreate(ctx, resp.ID, con2)
	if err != nil {
		panic(err)
	}
	fmt.Println("ExecProcess Created")
	cli.ContainerExecStart(ctx, execID2.ID, types.ExecStartCheck{})

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}
	if out == nil {

	}
	//fmt.Printf("%T", out)
	RDCloser, stpath, err := cli.CopyFromContainer(ctx, resp.ID, "/vol/data.txt")
	if err != nil {
		panic(err)
	}

	fmt.Print(RDCloser)
	fmt.Print(stpath)

	var timout time.Duration = 1

	buf := new(bytes.Buffer)

	buf.ReadFrom(out)

	st := buf.String()

	cli.ContainerStop(ctx, resp.ID, &timout)

	fmt.Print(st)

	io.Copy(os.Stdout, out)

}
