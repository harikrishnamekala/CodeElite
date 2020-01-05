package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

type OutputTeplStr struct {
	Code   string
	Output string
	Errors string
	Input  string
}

//To Stop the Running Container
func StopContainer(containerId string) bool {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	var timeout time.Duration = 3
	err = cli.ContainerStop(ctx, containerId, &timeout)
	if err != nil {
		return false
	}
	return true
}

//The Function will be called at hostenvcreator.go::CreateRespectiveEnvOfLanguage()
func CreateContainerExecEnv() string {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:           "compiler",
		Tty:             true,
		Cmd:             []string{"bash"},
		NetworkDisabled: true,
	}, nil, nil, "")
	if err != nil {
		log.Fatal("Error Occured While Creating a Container.")
		panic(err)
	}
	fmt.Print("Created Container : ")
	fmt.Print(resp.Warnings)
	fmt.Print(resp.ID + "\n")
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	fmt.Print("Started Container\n")

	return resp.ID

}
func StopContainerEnv(id string) {

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	var timout time.Duration = 1

	cli.ContainerStop(ctx, id, &timout)

	fmt.Println("Stopped the Container " + id + "\n")
}

/*
SOCK STREAM
The Function Copies the Files in Host which are located in Folder with Docker ID
 in the Host to Docker Contianer
*/
func CopyFilesToContainer(filenames []string, folderHostEnv string) {

	container_id := folderHostEnv

	for _, iter_filename := range filenames {
		_, copyfile := exec.Command("/bin/bash", "-c", "docker cp ./hostEnv/"+folderHostEnv+"/"+iter_filename+" "+container_id+":/home/").Output()
		if copyfile != nil {
			panic(copyfile)
		}
	}

}

func GetTheFilesFromContainer(filenames []string, container_id string) {
	for _, iter_filename := range filenames {
		_, getFileBack := exec.Command("/bin/bash", "-c", "docker cp "+container_id+":/home/"+iter_filename+" ./hostEnv/"+container_id+"/"+iter_filename).Output()
		if getFileBack != nil {
			panic(getFileBack)
		}
	}
}
func CreateOutputErrors(language_id int) {

}
func Runcode(language_id int, code string, input_for_code string) OutputTeplStr {

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	_, container_rep_id := CreateRespectiveEnvOfLanguage(language_id, "./hostEnv", "./language_templates")

	filenames := ReturnFilesList(language_id)

	writesourcefile := filenames[0]

	fmt.Print(writesourcefile)

	fmt.Print(code)

	fileSCode := []byte(code)

	fmt.Print(fileSCode)

	err = ioutil.WriteFile("./hostEnv/"+container_rep_id+"/"+writesourcefile, fileSCode, 0777)
	if err != nil {
		panic(err)
	}

	inputForCodeByte := []byte(input_for_code)

	err = ioutil.WriteFile("./hostEnv/"+container_rep_id+"/input.txt", inputForCodeByte, 0777)
	if err != nil {
		panic(err)
	}

	CopyFilesToContainer(filenames, container_rep_id)

	execution_st := types.ExecConfig{
		Cmd:          []string{"sh", "/home/compile.sh"},
		Tty:          true,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Detach:       true,
	}

	execution_ide, err := cli.ContainerExecCreate(ctx, container_rep_id, execution_st)
	if err != nil {
		panic(err)
	}

	err = cli.ContainerExecStart(ctx, execution_ide.ID, types.ExecStartCheck{
		Tty: true,
	})
	if err != nil {
		panic(err)
	}

	var filenametoget = []string{"errors.txt", "output.txt"}

	GetTheFilesFromContainer(filenametoget, container_rep_id)

	program_output, err := ioutil.ReadFile("./hostEnv/" + container_rep_id + "/output.txt")
	if err != nil {
		panic(err)
	}
	program_errors_output, err := ioutil.ReadFile("./hostEnv/" + container_rep_id + "/errors.txt")
	if err != nil {
		panic(err)
	}

	program_output_str := string(program_output)
	program_errors_output_str := string(program_errors_output)

	templateObjVal := OutputTeplStr{
		Code:   code,
		Output: program_output_str,
		Errors: program_errors_output_str,
		Input:  input_for_code,
	}

	go StopContainer(container_rep_id)

	return templateObjVal
}

/*

	con := types.ExecConfig{
		Cmd:          []string{"gcc", "/vol/main.c", "-o", "/vol/main", "2>", "/vol/errors.txt"},
		Tty:          true,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Detach:       true,
	}

	execID, err := cli.ContainerExecCreate(ctx, container_rep_id, con)
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

	execID3, err := cli.ContainerExecCreate(ctx, container_rep_id, types.ExecConfig{
		Cmd:          []string{"touch", "/vol/data.txt"},
		Tty:          true,
		AttachStdout: true,
		AttachStdin:  true,
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

	execID4, err := cli.ContainerExecCreate(ctx, container_rep_id, types.ExecConfig{
		Cmd:          []string{"touch", "/vol/errors.txt"},
		Tty:          true,
		AttachStdout: true,
		AttachStdin:  true,
		Detach:       true,
	})
	if err != nil {
		panic(err)
	}
	err = cli.ContainerExecStart(ctx, execID4.ID, types.ExecStartCheck{
		Tty: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created the File in the Container " + execID4.ID + "\n")

	/*copyfiletoBin, err := cli.ContainerExecCreate(ctx, resp.ID, types.ExecConfig{
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
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Detach:       true,
	}
	execID2, err := cli.ContainerExecCreate(ctx, container_rep_id, con2)
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

	out, err := cli.ContainerLogs(ctx, container_rep_id, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}
	if out == nil {

	}
	// TODO: Change the Folder
	_, getTheFileBack := exec.Command("/bin/bash", "-c", "docker cp "+container_rep_id+":/vol/data.txt ./controller/vol/"+randfolder+"/data.txt").Output()
	if getTheFileBack != nil {

	}
	_, getTheErrorsFileBack := exec.Command("/bin/bash", "-c", "docker cp "+container_rep_id+":/vol/errors.txt ./controller/vol/"+randfolder+"/errors.txt").Output()
	if getTheErrorsFileBack != nil {
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
	}
	//fmt.Print(st)

	//io.Copy(os.Stdout, out)

}
*/
