package main

//Importing Standard Packages
import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Print("Started Server at 8080")
	//http.Handle("/", http.FileServer(http.Dir("node_modules")))
	router := httprouter.New()
	router.GET("/", showIndex)
	//router.GET("/", http.FileServer(http.Dir(".")))
	router.GET("/execute/:code", executecode)
	//http.HandleFunc("/execute", executecode)
	http.ListenAndServe(":8080", router)

}

/*func FileServe(rw http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("views"))
}*/
func executecode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	code := []byte(ps.ByName("code"))
	err := ioutil.WriteFile("main.c", code, 777)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	//var options types.CopyToContainerOptions

	response, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "compiler",
	}, nil, nil, "")

	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, response.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	content := bytes.NewReader(code)
	if err := cli.CopyToContainer(ctx, response.ID, "/main.c", content, types.CopyToContainerOptions{}); err != nil {
		panic(err)
	}

	exec_id, err := cli.ContainerExecCreate(ctx, response.ID, types.ExecConfig{
		Cmd: []string{"gcc", "/main.c -o main"},
	})
	Use(exec_id)
	if err != nil {
		panic(err)
	}
	executable_id := "1"
	if err := cli.ContainerExecStart(ctx, executable_id, types.ExecStartCheck{}); err != nil {
		panic(err)
	}

	exect_id, err := cli.ContainerExecCreate(ctx, response.ID, types.ExecConfig{
		Cmd: []string{"./main"},
	})
	Use(exect_id)
	if err != nil {
		panic(err)
	}
	if err := cli.ContainerExecStart(ctx, executable_id, types.ExecStartCheck{}); err != nil {
		panic(err)
	}

	//fmt.Fprintf(w, "%s", ps.ByName("code"))
}
func showIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	static_html, err := ioutil.ReadFile("views/index.html")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", static_html)
}
func StopAllContainers() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
			panic(err)
		}
	}
}
func Use(something types.IDResponse) {

}
