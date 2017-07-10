package main

//Importing Standard Packages
import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"

	"codeelite.com/controller"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Print("Started Server at 8080")
	//http.Handle("/", http.FileServer(http.Dir("node_modules")))
	router := httprouter.New()
	//resourcesRouter()
	router.GET("/", showIndex)
	//router.GET("/", http.FileServer(http.Dir(".")))
	//router.ServeFiles("./node_modules", http.Dir("./node_modules"))
	router.GET("/executecode/:code", executecode)
	http.ListenAndServe(":8080", router)

}

/*func FileServe(rw http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("views"))
}*/
func executecode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	json_code := []byte(ps.ByName("code"))
	type Coder_code struct {
		codee []byte
	}
	var codes []Coder_code
	err := json.Unmarshal(json_code, &codes)
	if err != nil {
		panic(err)
	}
	code := codes[0].codee
	err = ioutil.WriteFile("./controller/vol/main.c", code, 0777)
	if err != nil {
		panic(err)
	}
	fmt.Print(code)

	go runner.Runcode()

	time.Sleep(time.Second * 10)
	//output, err := ioutil.ReadFile("./controller/vol/data.txt")]

	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", string(code))
	fmt.Fprintf(w, "%s", ps.ByName("code"))
}
func showIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	templ, err := template.ParseFiles("views/index.html")
	//static_html, err := ioutil.ReadFile("views/index.html")

	if err != nil {
		panic(err)
	}

	err = templ.Execute(w, nil)
	if err != nil {
		panic(err)
	}
	//fmt.Fprintf(w, "%s", static_html)

}

/*func resourcesRouter() {
	searchDir := "./node_modules"

	fileList := []string{}
	_ = filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})
	//router := httprouter.New()

	for _, file := range fileList[4:] {

		fmt.Println(file)

		//router.ServeFiles("./"+file, http.Dir(searchDir))
	}
	fmt.Print("------------------------------------------------------------Statin Glb")
	arr, err := filepath.Glob("node_modules/*")
	if err != nil {
		panic(err)
	}
	for _, data := range arr {
		fmt.Println(data)
		router.ServeFiles("./"+data, http.Dir(searchDir))
	}

}*/
