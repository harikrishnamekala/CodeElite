package runner

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func Runcode() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:   "compiler",
		Tty:     true,
		Cmd:     []string{"bash"},
		Volumes: map[string]struct{}{"./vol/": {}},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}
	fmt.Print(resp.Warnings)
	fmt.Print(resp.ID + "\n")
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	fmt.Print("Started Container\n")

	/*_, runit := exec.Command("/bin/bash", "-c", "docker start "+resp.ID[:12]).Output()
	if runit != nil {
		panic(runit)
	}*/

	_, copyfile := exec.Command("/bin/bash", "-c", "docker cp ./vol/main.c "+resp.ID+":/vol/").Output()
	if copyfile != nil {
		panic(copyfile)
	}
	fmt.Print("Copied File to container\n")

	_, copyfile2 := exec.Command("/bin/bash", "-c", "docker cp ./vol/compile.sh "+resp.ID+":/vol/").Output()
	if copyfile != nil {
		panic(copyfile2)
	}
	fmt.Print("Copied Shell to container\n")

	con := types.ExecConfig{
		Cmd:          []string{"gcc", "/vol/main.c", "-o", "/vol/main"},
		Tty:          true,
		AttachStdin:  false,
		AttachStdout: true,
		Detach:       true,
	}

	execID, err := cli.ContainerExecCreate(ctx, resp.ID, con)
	if err != nil {
		panic(err)
	}
	fmt.Println("ExecProcess Created" + execID.ID + "\n")
	err = cli.ContainerExecStart(ctx, execID.ID, types.ExecStartCheck{
		Tty: true,
	})
	if err != nil {
		panic(err)
	}

	execID3, err := cli.ContainerExecCreate(ctx, resp.ID, types.ExecConfig{
		Cmd:          []string{"touch", "/vol/data.txt"},
		Tty:          true,
		AttachStdout: true,
		AttachStdin:  false,
		Detach:       true,
	})
	if err != nil {
		panic(err)
	}
	err = cli.ContainerExecStart(ctx, execID3.ID, types.ExecStartCheck{
		Tty: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created the File in the Container " + execID3.ID + "\n")

	copyfiletoBin, err := cli.ContainerExecCreate(ctx, resp.ID, types.ExecConfig{
		Tty:          true,
		Cmd:          []string{"mv", "/vol/main", "/bin/main"},
		AttachStdin:  true,
		AttachStdout: true,
		Detach:       true,
	})
	if err != nil {
		panic(err)
	}
	err = cli.ContainerExecStart(ctx, copyfiletoBin.ID, types.ExecStartCheck{
		Tty: true,
	})
	fmt.Println("Moved File to Bin Folder")
	con2 := types.ExecConfig{
		Cmd:          []string{"sh", "/vol/compile.sh"},
		Tty:          true,
		AttachStdin:  false,
		AttachStdout: false,
		Detach:       true,
	}
	execID2, err := cli.ContainerExecCreate(ctx, resp.ID, con2)
	if err != nil {
		panic(err)
	}
	fmt.Println("ExecProcess Created" + execID2.ID + "\n")
	err = cli.ContainerExecStart(ctx, execID2.ID, types.ExecStartCheck{
		Tty: true,
	})
	if err != nil {
		panic(err)
	}
	//exec.Command("/bin/bash", "-c", "docker exec -t "+resp.ID+" /vol/main > /vol/data.txt")

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}
	if out == nil {

	}

	_, getTheFileBack := exec.Command("/bin/bash", "-c", "docker cp "+resp.ID+":/vol/data.txt ./vol/data.txt").Output()
	if getTheFileBack != nil {

	}
	//fmt.Printf("%T", out)
	/*RDCloser, _, err := cli.CopyFromContainer(ctx, resp.ID, "/vol/data.txt")
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)

	buf.ReadFrom(RDCloser)
	data := buf.Bytes()

	err = ioutil.WriteFile("./vol/data.txt", data, 0777)
	if err != nil {
		panic(err)
	}*/

	/*output := []byte(data)

	fmt.Println(output)

	err = ioutil.WriteFile("./vol/data.out", output, 0777)
	if err != nil {
		panic(err)
	}*/
	/*
			compile := exec.Command("/bin/bash", "-c", "docker exec "+resp.ID+" gcc /vol/main.c -o /vol/main")
			if compile != nil {

			}
			createoutputfile := exec.Command("/bin/bash", "-c", "docker exec "+resp.ID+" touch /vol/data.txt")

			if createoutputfile != nil {

			}

			execute := exec.Command("/bin/bash", "-c", "docker exec "+resp.ID+" /vol/main > data.txt")

			if execute != nil {

			}

		getTheFileBack := exec.Command("/bin/bash", "-c", "docker cp "+resp.ID+":/vol/data.txt /vol/data.txt")
		if getTheFileBack != nil {

		}*/

	/*getTheFileBack, statPath, err := cli.CopyFromContainer(ctx, resp.ID, "/vol/data.txt")
	if err != nil {
		panic(err)
	}*/
	var timout time.Duration = 1

	cli.ContainerStop(ctx, resp.ID, &timout)

	fmt.Println("Stopped the Container " + resp.ID + "\n")

	//fmt.Print(st)

	//io.Copy(os.Stdout, out)

}
